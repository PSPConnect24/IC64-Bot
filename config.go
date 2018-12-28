package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Config - Overall configuration
type Config struct {
	XMLName xml.Name  `xml:"config"`
	Bot     BotConfig `xml:"bot"`
}

// BotConfig - Bot-specific settings
type BotConfig struct {
	Token  string `xml:"token"`
	Prefix string `xml:"prefix"`
}

// Load - Loads the configuration into memory
func Load() Config {
	xmlFile, err := os.Open("config.xml")
	if err != nil {
		fmt.Println("An error occurred!\n", err)
	}
	defer xmlFile.Close()
	data, _ := ioutil.ReadAll(xmlFile)
	var config Config
	xml.Unmarshal(data, &config)
	return config
}
