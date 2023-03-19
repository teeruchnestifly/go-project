package services

import (
	"os"

	"go-rest-api/models/config"

	"gopkg.in/yaml.v3"
)

func LoadConfig() config.Config {
	configFile, err := os.Open("application.yaml")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config config.Config
	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}
