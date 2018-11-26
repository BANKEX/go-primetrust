package main

import (
	"github.com/BANKEX/go-primetrust"
	"github.com/BANKEX/go-primetrust/models"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	//accountId := os.Getenv("PRIMETRUST_ACCOUNT")
	webhookId := os.Getenv("PRIMETRUST_WEBHOOK")

	//webhook := models.NewWebhook(accountId, "http://www.ya1.ru", "test@test.ru","secret1234",true)
	//webhook.Data.ID = webhookId
	log.Printf("Getting webhook: %d", webhookId)
	var webhookPoint= &models.Webhook{}
	if webhookPoint, err := primetrust.GetWebhook(webhookId); err != nil {
		log.Println("Error getting webhook:", err)
	} else {
		log.Printf("Webhook: %+v", webhookPoint)
	}
	webhook:=*webhookPoint
	webhook.Data.Attributes.URL="http://ya1.ru"
	log.Printf("Updating webhook: %d", webhookId)
	if resp, err := primetrust.UpdateWebhook(webhookPoint); err != nil {
		log.Println("Error updating webhook:", err)
	} else {
		log.Println("Webhook updated OK")
		log.Panic(resp.Data)
	}
	log.Println("Done")
}
