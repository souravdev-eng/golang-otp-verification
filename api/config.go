package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(envName string) string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error while .env loading")
	}

	return os.Getenv(envName)
}
