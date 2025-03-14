package repository

import (
	"api/internal/models"
	"database/sql"
	"time"
)

type ClubRepository struct {
	db *sql.DB
}

func NewClubRepository(db *sql.DB) *ClubRepository {
	return &ClubRepository{
		db: db,
	}
}

func (r *ClubRepository) CreateClub(club models.Club) (string, error) {
	var clubID string
	err := r.db.QueryRow(`
		INSERT INTO clubs (name, description, email, member_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		club.Name,
		club.Description,
		club.Email,
		club.MemberCount,
		time.Now(),
		time.Now(),
	).Scan(&clubID)

	if err != nil {
		return "", err
	}

	return clubID, nil
}

func (r *ClubRepository) GetClubByID(clubID string) (*models.Club, error) {
	var club models.Club
	err := r.db.QueryRow(`
		SELECT id, name, description, email, member_count, created_at, updated_at
		FROM clubs
		WHERE id = $1`,
		clubID,
	).Scan(
		&club.ID,
		&club.Name,
		&club.Description,
		&club.Email,
		&club.MemberCount,
		&club.CreatedAt,
		&club.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &club, nil
}

func (r *ClubRepository) UpdateClub(club models.Club) error {
	_, err := r.db.Exec(`
		UPDATE clubs
		SET name = $1,
			description = $2,
			email = $3,
			member_count = $4,
			updated_at = $5
		WHERE id = $6`,
		club.Name,
		club.Description,
		club.Email,
		club.MemberCount,
		time.Now(),
		club.ID,
	)

	return err
}

func (r *ClubRepository) DeleteClub(clubID string) error {
	_, err := r.db.Exec("DELETE FROM clubs WHERE id = $1", clubID)
	return err
}

func (r *ClubRepository) ListClubs() ([]models.Club, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, email, member_count, created_at, updated_at
		FROM clubs
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clubs []models.Club
	for rows.Next() {
		var club models.Club
		err := rows.Scan(
			&club.ID,
			&club.Name,
			&club.Description,
			&club.Email,
			&club.MemberCount,
			&club.CreatedAt,
			&club.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		clubs = append(clubs, club)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clubs, nil
}
