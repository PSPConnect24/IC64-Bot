package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config represents the config
type Config struct {
	Token    string   `json:"token"`
	Prefixes []string `json:"prefixes"`
	Owners   []string `json:"owners"`
}

// Load the configuration into memory
func Load() Config {
	var config Config
	cfg, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Failed to open config.json: ", err)
		os.Exit(1)
	}
	err = json.NewDecoder(cfg).Decode(&config)
	if err != nil {
		fmt.Println("Failed to decode config.json: ", err)
		os.Exit(1)
	}
	cfg.Close()
	return config
}
