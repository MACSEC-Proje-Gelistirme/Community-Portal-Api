package repository

import (
	"database/sql"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user models.User) (sql.Result, error) {
	stmt := "INSERT INTO users (id, first_name, last_name, email, telephone_number, email_preferences, marketing_preferences, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	result, err := r.db.Exec(stmt, user.UserID, user.FirstName, user.LastName, user.Email, user.TelephoneNumber, user.EmailPreferences, user.MarketingPreferences, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return result, nil
}
