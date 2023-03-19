package services

import (
	"fmt"
	"os"

	"go-rest-api/models/config"

	"gopkg.in/yaml.v3"
)

func LoadConfig() {
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
	fmt.Println("hey")
	fmt.Println(config.App.Port)
}
