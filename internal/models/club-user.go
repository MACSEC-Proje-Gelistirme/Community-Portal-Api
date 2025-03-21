package models

type AddClubUserPayload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type DeleteClubUserPayload struct {
	UserID string `json:"user-id"`
}

type UpdateClubUserRolePayload struct {
	UserID string `json:"user-id"`
	Role   string `json:"role"`
}

type UserClubWithRole struct {
	ClubID      string `json:"club_id"`
	ClubName    string `json:"club_name"`
	Description string `json:"description"`
	Role        string `json:"role"`
}

type ClubMember struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type ClubDetailsResponse struct {
	Club    Club         `json:"club"`
	Members []ClubMember `json:"members"`
}
