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
	config           AWSConfig
	lambdaProperties []LambdaProperties
}

func NewAWS(config Configuration) AWS {
	return awss{
		config:           config.General.AWS,
		lambdaProperties: []LambdaProperties{},
	}
}

func (a awss) SearchRuntimeAllRegions() {
	for _, account := range a.config.Accounts {
		sess := a.generateAuth(account)
		regions := a.listAllRegions(sess)
		for _, region := range regions {
			a.runtimeSearchByIdentifier(sess, *region.RegionName, "")
		}
	}
}

func (a awss) SearchRuntimeByRegion(regionName string) {
	for _, account := range a.config.Accounts {
		sess := a.generateAuth(account)
		fmt.Println(a.runtimeSearchByIdentifier(sess, regionName, ""))
	}
}

func (a awss) runtimeSearchByIdentifier(sess *session.Session, regionName string, identifier string) []LambdaProperties {
	lambdas := a.lambdaProperties
	sess.Config.Region = aws.String(regionName)
	svcRegiao := lambda.New(sess)
	funcoes, err := svcRegiao.ListFunctions(nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, funcao := range funcoes.Functions {
		lambda := LambdaProperties{
			FunctionName: *funcao.FunctionName,
			FunctionARN:  *funcao.FunctionArn,
			Runtime:      *funcao.FunctionArn,
			Version:      *funcao.Version,
			LastModified: *funcao.LastModified,
		}
		lambdas = append(lambdas, lambda)
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
		Region:      aws.String("eu-west-1"),
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

	sess.Config.Credentials = credentials.NewStaticCredentials(*result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)
	return sess
}

func (a awss) readerMFA(account AWSAccount) string {
	fmt.Printf("Set MFA for %s: ", account.ARN)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(input, "\n")
}
