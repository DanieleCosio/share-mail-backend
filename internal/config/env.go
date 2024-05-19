package config

import (
	"log"
	"path"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	filePath := path.Join(AppConfig["ROOT_PATH"], "/.env")
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
