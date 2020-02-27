package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/igorhalfeld/latirebot/handles"
	"github.com/igorhalfeld/latirebot/services"
	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()

	scanHandler := handles.NewScanHandler(handles.Scan{
		TelegramService:  services.NewTelegramService(),
		RiachueloService: services.NewRiachueloService(),
		RennerService:    services.NewRennerService(),
	})

	scanHandler.Look()

	log.Println("Telegram KEY", os.Getenv("TELEGRAM_KEY"))

	c.AddFunc("@daily", scanHandler.Look)
	scanHandler.Look()

	go c.Start()

	log.Println("Bot is alive 🚨")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
