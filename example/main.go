package main

import (
	"log"

	"github.com/rbague/slack-webhook"
)

func main() {
	client := webhook.NewClient("WEBHOOK_URL_HERE")

	err := client.SendSimple("Hello, World!")
	if err != nil {
		log.Fatalf("could not send simple message: %v", err)
	}
}
