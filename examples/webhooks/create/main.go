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

	webhook := models.NewWebhook(accountId, "http://www.ya.ru", "test@test.ru","secret1234",true)

	if resp, err := primetrust.CreateNewWebhook(webhook); err != nil {
		log.Println("Error creating new webhook:", err)
	} else {
		log.Println("Webhook created OK")
		log.Panic(resp.Data)
	}
	log.Println("Done")
}
