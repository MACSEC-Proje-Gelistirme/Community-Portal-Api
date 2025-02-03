package repository

import (
	"database/sql"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetUserWithEmailPassword(payload models.GetUserWithEmailPasswordPayload) (sql.Result, error) {
	stmt := "SELECT * FROM users WHERE email = $1 AND password = $2"
	result, err := r.db.Exec(stmt, payload.Email, payload.Password)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CreateUser(payload models.CreateUserPayload) (sql.Result, error) {
	stmt := "INSERT INTO users (first_name, last_name, email, telephone_number, school_id, password) VALUES ($1, $2, $3, $4, $5, $6)"
	result, err := r.db.Exec(stmt, payload.FirstName, payload.LastName, payload.Email, payload.TelephoneNumber, payload.SchoolID, payload.Password)
	if err != nil {
		return nil, err
	}

	return result, nil
}
