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
	contactId := os.Getenv("PRIMETRUST_CONTACT")
	//file,_:= multipart.FileHeader{}("
	document := models.Document{
		AccountID:   accountId,
		ContactID:   contactId,
		Description: "Test upload for documentation",
		File:        nil,
		Extension:   ".jpg",
		Label:       "Test document",
		MimeType:    "image/jpeg",
		Public:      false,
	}

	if _, err := primetrust.UploadDocument(document); err != nil {
		log.Println("Error uploading new document:", err)
	} else {
		log.Println("Document uploaded OK")
	}

	log.Println("Done")
}
