package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/igorhalfeld/latirebot/handlers"
	"github.com/igorhalfeld/latirebot/repositories"
	"github.com/igorhalfeld/latirebot/services"
	"github.com/robfig/cron"
)

func createDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./latire.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createRepositoryContainer(db *sql.DB) repositories.Container {
	return repositories.Container{
		UserRepository: repositories.NewUserRepository(db),
	}
}

func createServiceContainer(repos repositories.Container) services.Container {
	return services.Container{
		UserService:      services.NewUserService(repos),
		RennerService:    services.NewRennerService(),
		RiachueloService: services.NewRiachueloService(),
		TelegramService:  services.NewTelegramService(),
	}
}

func createHandlerContainer(services services.Container) handlers.Container {
	return handlers.Container{
		HealthHandler:   handlers.NewHealthHandler(),
		ProductsHandler: handlers.NewProductHandler(services),
	}
}

func main() {
	telegramKey := os.Getenv("TELEGRAM_KEY")
	if telegramKey == "" {
		log.Println("Telegram KEY", "Not found")
	}

	db, err := createDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
	}

	c := cron.New()
	repos := createRepositoryContainer(db)
	services := createServiceContainer(repos)
	handlers := createHandlerContainer(services)

	http.HandleFunc("/", handlers.HealthHandler.Check)

	c.AddFunc("@daily", handlers.ProductsHandler.Look)
	//handlers.ProductsHandler.Look()
	go c.Start()

	log.Println("Bot is alive 🚨")
	http.ListenAndServe(":80", nil)
}
