package main

import (
	"github.com/BANKEX/go-primetrust"
	"github.com/BankEx/go-primetrust/models"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	accountId := os.Getenv("PRIMETRUST_ACCOUNT")
	webhookId := os.Getenv("PRIMETRUST_WEBHOOK")

	webhook := models.NewWebhook(accountId, "testUrl", true)
	webhook.Data.ID = webhookId
	if _, err := primetrust.UpdateWebhook(webhook); err != nil {
		log.Println("Error updating webhook:", err)
	} else {
		log.Println("Webhook updated OK")
	}
	log.Println("Done")
}
