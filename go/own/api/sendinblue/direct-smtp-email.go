package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SendinBlueMail struct {
	Sender      SendinBlueContact   `json:"sender"`
	To          []SendinBlueContact `json:"to"`
	HtmlContent string              `json:"htmlContent"`
	Subject     string              `json:"subject"`
}

type SendinBlueContact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

func main() {
	url := "https://api.sendinblue.com/v3/smtp/email"

	mail := SendinBlueMail{
		Sender: SendinBlueContact{
			Name:  "Николай Андреев",
			Email: "fake@gmail.com",
		},
		To: []SendinBlueContact{
			SendinBlueContact{
				Email: "myg0t@inbox.lv",
			},
		},
		HtmlContent: "<html><body><p>Привет, это тестовое письмо!</p></body></html>",
		Subject:     "Тестовое письмо",
	}

	jsonBody, _ := json.Marshal(mail)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("POST request creation error: ", err)
		return
	}

	apiKey := "...."

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request sending error: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response: ", resp.Status)
}
