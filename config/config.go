package config

import (
	"encoding/json"
	"fmt"
	"os"
)



var(
	Token string
	BotPrefix string

	config *Config
)

type Config struct {
	Token string `json:"TOKEN"`
	Prefix string `json:"BOT_PREFIX"`
}

func LoadConfig() error{
	fmt.Println("Loading config...")
	file, err := os.ReadFile("config/config.json")

	if err != nil {
		fmt.Println("Error while loading the file")
		return err
}

	fmt.Println("Converting the file to config struct")
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error while converting the file")
		return err
	}
	
	Token = config.Token
	BotPrefix = config.Prefix

	return nil

}



