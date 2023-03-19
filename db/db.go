package db

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

var DB *gorm.DB

func Connect() {
	configFile, err := os.Open("application.yaml")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config Config
	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		panic(err)
	}

	// Connect to database using GORM
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	// Set up database connection
	// dsn := "root:password@tcp(localhost:3306)/nestifly?charset=utf8mb4&parseTime=True&loc=Local"
	var err2 error
	DB, err2 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err2 != nil {
		fmt.Println(err)
		return
	}
}
