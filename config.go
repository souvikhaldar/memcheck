package main

import (
	"encoding/json"
	"log"
	"os"

	"io/ioutil"
)

type config struct {
	SourceMail     string
	SourcePassword string
	TargetMail     []string
}

func ParseConfig() *config {
	conf := &config{}
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Println("Unable to read config file:", err)
		return nil
	}
	jsonInp, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("Unable to read: ", err)
		return conf
	}
	if err := json.Unmarshal(jsonInp, conf); err != nil {
		log.Println("Unable to parse JSON: ", err)
		return conf
	}
	return conf
}
