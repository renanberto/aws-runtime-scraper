package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Configuration

func init() {
	file, err := os.ReadFile("config/general.yml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}
}
