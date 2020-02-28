package main

import (
	"log"
	"os"

	"github.com/igorhalfeld/latirebot/handles"
	"github.com/igorhalfeld/latirebot/services"
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()

	healthHandler := handles.NewHealthHandler()
	scanHandler := handles.NewScanHandler(handles.Scan{
		TelegramService:  services.NewTelegramService(),
		RiachueloService: services.NewRiachueloService(),
		RennerService:    services.NewRennerService(),
	})

	log.Println("Telegram KEY", os.Getenv("TELEGRAM_KEY"))

	c.AddFunc("@daily", scanHandler.Look)
	go c.Start()

	log.Println("Bot is alive 🚨")

	scanHandler.Look()
	healthHandler.Check()
}
