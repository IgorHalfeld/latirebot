package services

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/igorhalfeld/latirebot/structs"
)

// TelegramService model
type TelegramService struct{}

// NewTelegramService creates a new instance
func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

// Send is the trigger to send messages
func (t *TelegramService) Send(p structs.SendPayload) {
	params := "?photo=" + p.Photo + "&chat_id=" + p.ChatID + "&parse_mode=" + p.ParseMode + "&caption=" + p.Caption
	token := os.Getenv("TELEGRAM_KEY")
	URL := "https://api.telegram.org/bot" + token + "/sendPhoto" + params

	request, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Sent! ", p.Name)
	defer response.Body.Close()
}
