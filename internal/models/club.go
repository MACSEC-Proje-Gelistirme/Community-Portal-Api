package models

type Club struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	MemberCount string `json:"member_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ClubWithRole struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type CreateClubPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
}
