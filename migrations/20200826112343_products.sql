-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
	id              TEXT PRIMARY KEY UNIQUE,
  sku             TEXT,
	name            TEXT NOT NULL,
  provider        TEXT NOT NULL,
  response_data   TEXT NOT NULL,
	normal_price    FLOAT NOT NULL,
	discount_price  FLOAT NOT NULL,
  created_at      TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
