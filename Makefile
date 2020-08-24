migration-up:
	@goose --dir migrations sqlite3 latire.db up
