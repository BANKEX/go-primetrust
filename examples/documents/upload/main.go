package main

import (
	"github.com/BANKEX/go-primetrust"
	"log"
	"os"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	accountId := "6db32598-6ccc-452d-98d6-f1338f982181"
	contactId := "8d60c332-ab08-4ee3-815c-169e45ef4c09"

	path, _ := os.Getwd()
	path += "/examples/documents/upload/empty.jpg"
	log.Println(path)

	if response, err := primetrust.UploadDocument(path, accountId,contactId,"description","extention","label","mimeType"); err != nil {
		log.Println("Error uploading new document:", err)
	} else {
		log.Println("Document uploaded: ", response)
	}

	log.Println("Done")
}