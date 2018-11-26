package primetrust

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BANKEX/go-primetrust/models"
	"io/ioutil"
	"net/http"
	"log"
)

func UploadDocument(document models.Document) (*models.DocumentResponse, error) {
	apiUrl := fmt.Sprintf("%s/uploaded-documents", _apiPrefix)

	jsonBytes, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	buffer:=bytes.NewBuffer(jsonBytes)
	log.Println(buffer)
	req, err := http.NewRequest("POST", apiUrl, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "multipart/form-data")
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
