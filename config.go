package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Token    string   `json:"token"`
	Prefixes []string `json:"prefixes"`
	Owners   []string `json:"owners"`
	DSN      string   `json:"dsn"`
}

func load() config {
	var conf config
	cfg, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Failed to open config.json: ", err)
		os.Exit(1)
	}
	err = json.NewDecoder(cfg).Decode(&conf)
	if err != nil {
		fmt.Println("Failed to decode config.json: ", err)
		os.Exit(1)
	}
	cfg.Close()
	return conf
}
