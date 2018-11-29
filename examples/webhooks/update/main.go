package main

import (
	"github.com/BANKEX/go-primetrust"
	"github.com/BANKEX/go-primetrust/models"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	accountId := os.Getenv("PRIMETRUST_ACCOUNT")
	webhookId := os.Getenv("PRIMETRUST_WEBHOOK")

	log.Printf("Getting webhook: %d", webhookId)

	webhook := models.NewWebhook(accountId, "http://www.ya1.ru", "test@test.ru", "secret1234", true)
	webhook.Data.ID = webhookId
	if webhook, err := primetrust.UpdateWebhook(webhook); err != nil {
		log.Println("Error updating webhook:", err)
	} else {
		log.Println("Webhook updated OK")
		log.Panic(webhook.Data)
	}
	log.Println("Done")
}
