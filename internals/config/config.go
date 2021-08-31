package config

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func ReadConfig(configPath string) config {
	updateConfig := func(filePath string, config *config) {
		configFile, err := os.Open(filePath)
		defer configFile.Close()

		if err != nil {
			log.Fatal(err)
		}

		jsonDecoder := json.NewDecoder(configFile)
		jsonDecoder.Decode(&config)
	}

	var config config
	updateConfig(configPath, &config)

	return config
}
