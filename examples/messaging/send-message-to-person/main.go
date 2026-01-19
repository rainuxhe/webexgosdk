package main

import (
	"context"
	"log"

	"github.com/rainuxhe/webexgosdk"
	"github.com/rainuxhe/webexgosdk/messaging"
)

func main() {
	accessToken := ""
	toPersonEmail := ""
	client, err := webexgosdk.NewClient(accessToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	message, err := client.Messaging.Messages.Create(ctx, &messaging.MessageCreateRequest{
		ToPersonEmail: toPersonEmail,
		Markdown:      "This is message from webexgossdk",
	})

	if err != nil {
		log.Fatalf("Failed to create message: %v", err)
	}
	log.Printf("Message created: %v", message)
}
