package components

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var Configuration = ConfigurationStruct{
	Port:    "8080",
	Cookies: "",
}

type ConfigurationStruct struct {
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
			err = decoder.Decode(&Configuration)
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
		err = encoder.Encode(Configuration)
	}

	return err
}

func QuickSaveConfig() {
	if err := SaveConfig(); err != nil {
		log.Println(err)
	}
}
