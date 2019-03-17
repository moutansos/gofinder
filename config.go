package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	BotKey           string `json:"botKey"`
	ChannelId        string `json:"channelId"`
	ConnectionString string `json:"connectionString"`
}

func readConfig(file string) Config {
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	var config Config
	err = json.Unmarshal(buff, &config)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}

	return config
}
