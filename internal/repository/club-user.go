package repository

import (
	"api/internal/models"
	"database/sql"
	"time"
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

func (c *ClubUserRepository) GetClubsWithUserID(userID string) ([]models.ClubWithRole, error) {
	var clubs []models.ClubWithRole

	rows, err := c.db.Query(`
		SELECT c.id, c.name, c.description, c.email, cr.role
		FROM clubs c
		JOIN club_roles cr ON c.id = cr.club_id
		WHERE cr.user_id = $1 AND cr.role != 'member';`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var club models.ClubWithRole
		err := rows.Scan(
			&club.ID,
			&club.Name,
			&club.Description,
			&club.Email,
			&club.Role,
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

func (c *ClubUserRepository) CreateClubRole(clubID string, userID string, role string) (bool, error) {
	_, err := c.db.Exec(`
		INSERT INTO club_roles (user_id, club_id, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)`,
		userID, clubID, role, time.Now(),
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *ClubUserRepository) DeleteAllClubRoles(clubID string) error {
	stmt := `
		DELETE FROM club_roles WHERE club_id = $1 
	`
	_, err := c.db.Exec(stmt, clubID)

	if err != nil {
		return err
	}

	return nil
}

func (c *ClubUserRepository) DeleteClubRole(clubID string, userID string) error {
	_, err := c.db.Exec(`
		DELETE FROM club_roles 
		WHERE club_id = $1 AND user_id = $2`,
		clubID, userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClubUserRepository) UpdateClubRole(clubID string, userID string, role string) error {
	_, err := c.db.Exec(`
		UPDATE club_roles 
		SET role = $3, updated_at = $4
		WHERE club_id = $1 AND user_id = $2`,
		clubID, userID, role, time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClubUserRepository) GetUserClubsWithRoles(userID string) ([]models.UserClubWithRole, error) {
	var clubs []models.UserClubWithRole

	rows, err := c.db.Query(`
		SELECT c.id, c.name, c.description, cr.role
		FROM clubs c
		JOIN club_roles cr ON c.id = cr.club_id
		WHERE cr.user_id = $1;`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var club models.UserClubWithRole
		err := rows.Scan(
			&club.ClubID,
			&club.ClubName,
			&club.Description,
			&club.Role,
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

func (c *ClubUserRepository) GetClubDetailsWithMembers(clubID string) (*models.Club, []models.ClubMember, error) {
	// Get club details
	var club models.Club
	err := c.db.QueryRow(`
		SELECT id, name, description, email, member_count, created_at, updated_at
		FROM clubs
		WHERE id = $1`, clubID).Scan(
		&club.ID,
		&club.Name,
		&club.Description,
		&club.Email,
		&club.MemberCount,
		&club.CreatedAt,
		&club.UpdatedAt,
	)
	if err != nil {
		return nil, nil, err
	}

	// Get club members
	rows, err := c.db.Query(`
		SELECT u.id, u.first_name, u.last_name, u.email, cr.role
		FROM users u
		JOIN club_roles cr ON u.id = cr.user_id
		WHERE cr.club_id = $1`, clubID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var members []models.ClubMember
	for rows.Next() {
		var member models.ClubMember
		err := rows.Scan(
			&member.UserID,
			&member.FirstName,
			&member.LastName,
			&member.Email,
			&member.Role,
		)
		if err != nil {
			return nil, nil, err
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return &club, members, nil
}
