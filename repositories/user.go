package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/igorhalfeld/latirebot/structs"
	"github.com/jmoiron/sqlx"
)

var (
	errorCreate  = errors.New("User create failed")
	errorReadAll = errors.New("User readall failed")
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) Create(ctx context.Context, user structs.User) error {
	query := `
		INSERT INTO
			users(id, telegram_id, name, username, clothing_type, started_at)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`
	result, err := ur.db.ExecContext(ctx, query, user.ID, user.TelegramID, user.Name, user.Username, user.ClothingType, user.StartedAt)
	if err != nil {
		log.Println(err)
		return errorCreate
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Println(errorCreate, err)
		return errorCreate
	}
	log.Println("affected rows", affectedRows)

	return nil
}

func (ur UserRepository) ReadAll(ctx context.Context) ([]structs.User, error) {
	var users []structs.User
	query := `
		SELECT * FROM users LIMIT 50
	`

	err := ur.db.SelectContext(ctx, &users, query)
	if err != nil {
		log.Fatalln(errorReadAll, err)
		return []structs.User{}, errorReadAll
	}

	return users, nil
}

func (ur UserRepository) ReadOneByUsername(context context.Context) (structs.User, error) {
	return structs.User{}, nil
}
