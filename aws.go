package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/sts"
)

type awss struct {
	scrapper         SScraper
	config           AWSConfig
	lambdaProperties []LambdaProperties
}

func NewAWS(config Configuration, scrapper SScraper) AWS {
	return awss{
		scrapper:         scrapper,
		config:           config.General.AWS,
		lambdaProperties: []LambdaProperties{},
	}
}

func (a awss) SearchRuntimeAllRegions() {
	for _, account := range a.config.Accounts {
		sess := a.generateAuth(account)
		regions := a.listAllRegions(sess)
		for _, region := range regions {
			for _, table := range a.scrapper.table {
				lambdas := a.runtimeSearchByIdentifier(sess, *region.RegionName, table.Identifier)
				for _, v := range lambdas {
					fmt.Println(v.FunctionName)
				}
			}
		}
	}
}

func (a awss) SearchRuntimeByRegion(regionName string) {
	for _, account := range a.config.Accounts {
		sess := a.generateAuth(account)
		for _, table := range a.scrapper.table {
			lambdas := a.runtimeSearchByIdentifier(sess, regionName, table.Identifier)
			for _, v := range lambdas {
				fmt.Println(v.FunctionName)
			}
		}
	}
}

func (a awss) runtimeSearchByIdentifier(sess *session.Session, regionName string, identifier string) []LambdaProperties {
	lambdas := a.lambdaProperties
	sess.Config.Region = aws.String(regionName)
	service := lambda.New(sess)

	list, err := service.ListFunctions(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range list.Functions {
		if *f.Runtime == identifier {
			lambda := LambdaProperties{
				FunctionName: *f.FunctionName,
				FunctionARN:  *f.FunctionArn,
				Runtime:      *f.Runtime,
				Version:      *f.Version,
				LastModified: *f.LastModified,
			}
			lambdas = append(lambdas, lambda)
		}
	}
	return lambdas
}

func (a awss) listAllRegions(sess *session.Session) []*ec2.Region {
	result, err := ec2.New(sess).DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		log.Fatal(err)
	}
	return result.Regions
}

func (a awss) generateAuth(account AWSAccount) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"), // default login
		Credentials: credentials.NewStaticCredentials(account.Token, account.Key, ""),
	})

	if err != nil {
		log.Fatal(err)
	}

	mfa := a.readerMFA(account)

	result, err := sts.New(sess).GetSessionToken(&sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
		SerialNumber:    aws.String(account.ARN),
		TokenCode:       aws.String(mfa),
	})

	if err != nil {
		log.Fatal(err)
	}

	sess.Config.Credentials = credentials.NewStaticCredentials(
		*result.Credentials.AccessKeyId,
		*result.Credentials.SecretAccessKey,
		*result.Credentials.SessionToken,
	)
	return sess
}

// MFA Reader :(
func (a awss) readerMFA(account AWSAccount) string {
	fmt.Printf("Set MFA for %s: ", account.ARN)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(input, "\n")
}
