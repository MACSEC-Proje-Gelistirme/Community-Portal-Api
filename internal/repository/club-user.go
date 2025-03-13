package repository

import (
	"api/internal/models"
	"database/sql"
)

type ClubUserRepository struct {
	db *sql.DB
}

func NewClubUserRepository(db *sql.DB) *ClubUserRepository {
	return &ClubUserRepository{
		db: db,
	}
}

func (c *ClubUserRepository) GetUserRole(clubID, userID string) (string, error) {
	var role string
	err := c.db.QueryRow("SELECT role FROM club_roles WHERE club_id = $1 AND user_id = $2", clubID, userID).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (c *ClubUserRepository) GetClubsWithUserID(userID string) ([]models.Club, error) {
	var clubs []models.Club

	rows, err := c.db.Query(`
		SELECT c.id, c.name, c.description, c.email, c.member_count, c.created_at, c.updated_at
		FROM clubs c
		JOIN club_roles cr ON c.id = cr.club_id
		WHERE cr.user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
