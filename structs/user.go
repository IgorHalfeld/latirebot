package structs

import (
	"github.com/google/uuid"
)

type ClothingEnum string

const (
	ClothingTypeMale   ClothingEnum = "MALE"
	ClothingTypeFemale ClothingEnum = "FEMALE"
	ClothingTypeBoth   ClothingEnum = "BOTH"
)

type User struct {
	ID           uuid.UUID    `db:"id"`
	TelegramID   int64        `db:"telegram_id"`
	Name         string       `db:"name"`
	Username     string       `db:"username"`
	StartedAt    string       `db:"started_at"`
	ClothingType ClothingEnum `db:"clothing_type"`
}
