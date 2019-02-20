package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Token    string   `yaml:"token"`
	Prefixes []string `yaml:"prefixes"`
	Owners   []string `yaml:"owners"`
}

func load() config {
	var conf config
	cfg, err := os.Open("config.yml")
	if err != nil {
		fmt.Println("Failed to open config.yml: ", err)
		os.Exit(1)
	}
	err = yaml.NewDecoder(cfg).Decode(&conf)
	if err != nil {
		fmt.Println("Failed to decode config.yml: ", err)
		os.Exit(1)
	}
	cfg.Close()
	return conf
}
