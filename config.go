package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var config Configuration

func init() {
	file, err := ioutil.ReadFile("config/general.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}
}
