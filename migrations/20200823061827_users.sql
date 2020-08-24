-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id              TEXT PRIMARY KEY UNIQUE,
	telegram_id     INTEGER UNIQUE,
	name            TEXT NOT NULL,
	username        TEXT NOT NULL UNIQUE,
	clothing_type   TEXT NOT NULL,
  started_at      TEXT 
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
