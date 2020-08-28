package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/igorhalfeld/latirebot/handlers"
	"github.com/igorhalfeld/latirebot/repositories"
	"github.com/igorhalfeld/latirebot/services"
	"github.com/robfig/cron"
)

func createDatabaseConnection() (*sqlx.DB, error) {
	env := os.Getenv("APP_ENV")
	db, err := sqlx.Open("sqlite3", "./latire-"+env+".db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createRepositoryContainer(db *sqlx.DB) repositories.Container {
	return repositories.Container{
		UserRepository:    repositories.NewUserRepository(db),
		ProductRepository: repositories.NewProductRepository(db),
	}
}

func createServiceContainer(repos repositories.Container) services.Container {
	telegramKey := os.Getenv("TELEGRAM_KEY")
	if telegramKey == "" {
		log.Println("Telegram KEY", "Not found")
	}

	return services.Container{
		RennerService:    services.NewRennerService(),
		RiachueloService: services.NewRiachueloService(),
		UserService:      services.NewUserService(repos),
		ProductService:   services.NewProductService(repos),
		TelegramService:  services.NewTelegramService(repos, telegramKey),
	}
}

func createHandlerContainer(services services.Container) handlers.Container {
	return handlers.Container{
		HealthHandler:   handlers.NewHealthHandler(),
		ProductsHandler: handlers.NewProductHandler(services),
	}
}

func main() {
	db, err := createDatabaseConnection()
	if err != nil {
		log.Fatalln(err)
	}

	c := cron.New()
	repos := createRepositoryContainer(db)
	services := createServiceContainer(repos)
	handlers := createHandlerContainer(services)

	handlers.ProductsHandler.Look()

	http.HandleFunc("/", handlers.HealthHandler.Check)
	c.AddFunc("@weekly", handlers.ProductsHandler.Look)

	go c.Start()
	go services.TelegramService.ListenMessages()

	log.Println("Bot is alive 🚨")
	http.ListenAndServe(":80", nil)
}
