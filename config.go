package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var configuration = Configuration{
	Port:    "8080",
	Cookies: "",
}

type Configuration struct {
	Port    string `json:"port"`
	Cookies string `json:"cookies"`
}

func LoadConfig() error {
	err := errors.New("")
	if _, err = os.Stat("./config.json"); err == nil {
		file, error := os.Open("./config.json")
		err = error
		if error == nil {
			defer file.Close()
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&configuration)
		}
	} else if os.IsNotExist(err) {
		err = SaveConfig()
	}
	return err
}

func SaveConfig() error {
	file, err := os.Create("config.json")
	if err == nil {
		defer file.Close()

		encoder := json.NewEncoder(file)
		err = encoder.Encode(configuration)
	}

	return err
}

func QuickSaveConfig() {
	if err := SaveConfig(); err != nil {
		log.Println(err)
	}
}
