package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Missing Api Key")
	}
}
