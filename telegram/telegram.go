package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
)


func SendMessage(text string, config *ini.File) {
	botApiToken := config.Section("telegram").Key("botApiToken").String()	
	chatId := config.Section("telegram").Key("chatId").String()

	// Compose a body of the request to the Telegram API
	body, _ := json.Marshal(map[string]string{
		"chat_id":              chatId,
		"text":                 text,
		"disable_notification": "true",
	})
	requestBody := bytes.NewBuffer(body)

	// Send a message via HTTP POST
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botApiToken)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatalf("ERROR: '%v'", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(responseBody)
	log.Println(sb)
}
