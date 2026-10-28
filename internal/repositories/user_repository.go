package repositories

import (
	"backend/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

var db *sql.DB

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository {
		db: db,
	}
}

func (r* UserRepository) Create(user models.User) error {
	query := `INSERT INTO users (email, password_hash) VALUES (?, ?)`

	_, err := r.db.Exec(query, user.Email, user.PasswordHash)

	return err
}