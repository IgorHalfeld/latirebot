package repositories

import (
	"context"
	"database/sql"

	"github.com/igorhalfeld/latirebot/structs"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) Create(context context.Context, ID string) error {
	return nil
}

func (ur UserRepository) ReadAll(context context.Context) ([]structs.Users, error) {
	return []structs.Users{}, nil
}
