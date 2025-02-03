package repository

import "database/sql"

type ClubUserRolesRepository struct {
	db *sql.DB
}

func NewClubUserRolesRepository(db *sql.DB) *ClubUserRolesRepository {
	return &ClubUserRolesRepository{
		db: db,
	}
}

func (r *ClubUserRolesRepository) GetUserRole(clubID, userID string) (string, error) {
	var role string
	err := r.db.QueryRow("SELECT role FROM club_roles WHERE club_id = $1 AND user_id = $2", clubID, userID).Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
}
