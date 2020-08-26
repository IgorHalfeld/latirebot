package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/google/uuid"
	"github.com/igorhalfeld/latirebot/repositories"
	"github.com/igorhalfeld/latirebot/structs"
)

const (
	maleClothing   string = "Roupas masculinas"
	femaleClothing string = "Roupas femininas"
	bothClothing   string = "Ambas"
)

type TelegramService struct {
	repos repositories.Container
	bot   *tgbotapi.BotAPI
}

func NewTelegramService(repos repositories.Container, token string) *TelegramService {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalln(err)
	}
	bot.Debug = true

	return &TelegramService{repos, bot}
}

func (ts TelegramService) ListenMessages() {
	ctx := context.Background()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := ts.bot.GetUpdatesChan(u)
	if err != nil {
		log.Println("fail on error updates")
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		user := structs.User{
			ID:         uuid.New(),
			Name:       update.Message.From.FirstName + " " + update.Message.From.LastName,
			Username:   update.Message.From.UserName,
			StartedAt:  time.Now().UTC().String(),
			TelegramID: update.Message.Chat.ID,
		}

		var clothingType structs.ClothingEnum

		switch update.Message.Text {
		case "/start":
			ts.handleWellcomeMessageAndSetup(user)
		case maleClothing:
			clothingType = "MALE"
		case femaleClothing:
			clothingType = "FEMALE"
		case bothClothing:
			clothingType = "BOTH"
		default:
			log.Println("No case for message:", update.Message.Text)
		}

		if clothingType != "" {
			user.ClothingType = clothingType
			err := ts.repos.UserRepository.Create(ctx, user)
			if err != nil {
				log.Panicln("error on create user", err)
			}
			ts.handleSetupDone(user)
		}
	}
}

func (ts TelegramService) SendNotification(payload structs.NotificationPayload) {
	product := payload.Product
	user := payload.User
	np := payload.NormalPrice
	dp := payload.DiscountPrice

	text := `<strong>` + strings.ToUpper(product.Name) + `</strong> ⚡️ ` + strings.ToUpper(string(product.Provider)) + ` ⚡️ ` +
		`<i>R$` + fmt.Sprintf("%.2f", dp) + `</i> - <s>R$` + fmt.Sprintf("%.2f", np) + `</s> ⚡️ ` +
		`<a href='` + product.Link + `'>Ver detalhes 💵</a>`

	msg := tgbotapi.NewPhotoShare(user.TelegramID, product.ImageURL)
	msg.ParseMode = "html"
	msg.Caption = text
	msg.ReplyMarkup = tgbotapi.PhotoConfig{
		Caption:   payload.Caption,
		ParseMode: "html",
	}
	ts.bot.Send(msg)
}

func (ts TelegramService) handleWellcomeMessageAndSetup(user structs.User) {
	text := `
<strong>Au Au Au, ` + user.Name + `! Eu sou o Latire</strong> 😍
Vou mandar para você os melhores descontos!!
Antes de tudo me fale quais <strong>tipos de roupas</strong> você quer receber descontos 🔥
	`
	msgWellcome := tgbotapi.NewMessage(user.TelegramID, text)
	msgWellcome.ParseMode = "html"
	ts.bot.Send(msgWellcome)

	labels := []tgbotapi.KeyboardButton{
		{
			Text:            maleClothing,
			RequestContact:  false,
			RequestLocation: false,
		},
		{
			Text:            femaleClothing,
			RequestContact:  false,
			RequestLocation: false,
		},
		{
			Text:            bothClothing,
			RequestContact:  false,
			RequestLocation: false,
		},
	}

	msgCouthingType := tgbotapi.NewMessage(user.TelegramID, "Selecione entre roupas masculinas, femininas ou ambas das duas!")
	msgCouthingType.ReplyMarkup = &tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        [][]tgbotapi.KeyboardButton{labels},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
		Selective:       true,
	}
	ts.bot.Send(msgCouthingType)
}

func (ts TelegramService) handleSetupDone(user structs.User) {
	text := `
<strong>Au Au Au, ` + user.Name + `! Tudo pronto!</strong> 🚀
Agora é só esperar, vou mandar os melhores descontos das lojas <strong>toda semana</strong>
`
	msgWellcome := tgbotapi.NewMessage(user.TelegramID, text)
	msgWellcome.ParseMode = "html"
	ts.bot.Send(msgWellcome)
}
