package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Configuration

func configure(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatal(err)
	}
}
