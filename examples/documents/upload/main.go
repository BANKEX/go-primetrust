package main

import (
	"github.com/BANKEX/go-primetrust"
	"log"
	"os"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	primetrust.Init(true, os.Getenv("PRIMETRUST_LOGIN"), os.Getenv("PRIMETRUST_PASSWORD"))

	r := mux.NewRouter()

	r.HandleFunc("/upload", UploadDocumentHandler)

	panic(http.ListenAndServe(":7000",r))
}

func UploadDocumentHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(w, r)

	r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "File upload error. Please ensure that size of image file is less than 5 MB", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	accountId := os.Getenv("PRIMETRUST_ACCOUNT")
	contactId := os.Getenv("PRIMETRUST_CONTACT")

	if err != nil {
		log.Println(err)
		http.Error(w, "Wrong issue date format. Please use YYYY-MM-DD (eg. 2010-12-25)", http.StatusInternalServerError)
		return
	}

	log.Println("Send request")
	if response, err := primetrust.UploadDocument(file, fileHeader.Filename, accountId,contactId,r.PostForm.Get("description"),r.PostForm.Get("extention"),r.PostForm.Get("label"),r.PostForm.Get("mimeType")); err != nil {
		log.Println("Error uploading new document:", err)
		http.Error(w, "Error uploading new document:", http.StatusInternalServerError)
	} else {
		log.Println("Document uploaded: ", response)
		w.Write([]byte("OK"))
	}

	return
}