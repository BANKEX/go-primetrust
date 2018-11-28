package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BANKEX/go-primetrust/models"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func UploadDocument(document models.Document) (*models.DocumentResponse, error) {
	apiUrl := fmt.Sprintf("%s/uploaded-documents", _apiPrefix)

	jsonBytes, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", document.Label)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, document.File); err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"account-id":  document.AccountID,
		"contact-id":  document.ContactID,
		"description": document.Description,
		"extension":   document.Extension,
		"label":       document.Label,
		"mime_type":   document.MimeType,
		"public":      document.Public,
	}

	for key, val := range data {
		valString, ok := val.(string)

		if !ok {
			return nil, err
		}
		if fw, err = w.CreateFormField(key); err != nil {
			return nil, err
		}
		if _, err = fw.Write([]byte(valString)); err != nil {
			return nil, err
		}
	}

	w.Close()

	buffer := bytes.NewBuffer(jsonBytes)
	log.Println(buffer)
	req, err := http.NewRequest("POST", apiUrl, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Add("Authorization", _authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, _ := ioutil.ReadAll(res.Body)

	response := models.DocumentResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, errors.New("Unmarshal error")
	}

	return &response, nil
}
