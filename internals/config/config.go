package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type config struct {
	SleepTimeout uint `json:"sleepTimeout"`
}

func ReadConfig(configPath string) {
	updateConfig := func(filePath string, config *config) {
		configFile, err := os.Open(filePath)
		defer func() {
			configFile.Close()
			fmt.Println("Файл '" + filePath + "' закрыт")
			time.Sleep(time.Second * time.Duration(config.SleepTimeout))
		}()

		if err != nil {
			fmt.Println(err)
			return
		}

		jsonDecoder := json.NewDecoder(configFile)
		jsonDecoder.Decode(&config)
	}

	var config config
	for {
		updateConfig(configPath, &config)
		fmt.Println("SleepTimeout ==", config.SleepTimeout)
	}
}
