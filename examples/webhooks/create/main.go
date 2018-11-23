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

	webhook := models.NewWebhook(accountId, "testUrl", true)

	if _, err := primetrust.CreateNewWebhook(webhook); err != nil {
		log.Println("Error uploading new document:", err)
	} else {
		log.Println("Document uploaded OK")
	}
}
