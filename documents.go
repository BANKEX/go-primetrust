package primetrust

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func UploadDocument(file multipart.File, filename string, accountId string, contactId string, description string, extension string, label string, mimeType string) (*map[string]interface{}, error) {
	apiUrl := fmt.Sprintf("%s/uploaded-documents", _apiPrefix)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"account-id":  accountId,
		"contact-id":  contactId,
		"description": description,
		"extension":   extension,
		"label":       label,
		"mime_type":   mimeType,
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

	req, err := http.NewRequest("POST", apiUrl, &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Add("Authorization", _authHeader)

	log.Println("make request")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	log.Printf("result  ", res)
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	body, _ := ioutil.ReadAll(res.Body)

	var resData map[string]interface{}
	log.Printf("result  ", body)

	if err := json.Unmarshal(body, &resData); err != nil {
		return nil, err
	}
	return &resData, nil
}
