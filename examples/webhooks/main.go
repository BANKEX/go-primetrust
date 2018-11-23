package main

import (
	"github.com/BANKEX/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))
	if webhooks, err := primetrust.GetWebhooks(); err != nil {
		log.Println("Error getting webhooks:", err.Error())
	} else {
		log.Printf("Webhooks: %d", len(webhooks.Data))
	}

	log.Println("Done")
}
