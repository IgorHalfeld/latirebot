package main

import (
	"fmt"
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

	c.AddFunc("@daily", scanHandler.Look)

	go c.Start()

	fmt.Println("Bot is alive 🚨")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
