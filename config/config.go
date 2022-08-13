package config

import (
	"log"

	"github.com/joho/godotenv"
)

var Configs map[string]string

func ReadConfig() {
	var err error
	if Configs, err = godotenv.Read(); err != nil {
		log.Println(".env file not found")
	}
	return
}
