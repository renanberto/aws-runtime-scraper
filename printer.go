package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func pretty(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func LambdaPrinter(lambdaProperties []LambdaProperties, region string) {
	var buffer bytes.Buffer
	err := pretty(lambdaProperties, &buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("########################## Region: %s \n", region)
	fmt.Println(buffer.String())
}
