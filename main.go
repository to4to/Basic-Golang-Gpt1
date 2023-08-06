package main

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Missing Api Key")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)


	response,err:=client.Completion(ctx,gpt3.CompletionRequest{})



	if err!=nil{
		log.Fatal(err)
	}
}
