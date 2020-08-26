build:
	go build -v main.go

migration-up:
	@goose --dir migrations sqlite3 latire.db up

run:
	TELEGRAM_KEY=<MY_KEY>
	go run main.go
