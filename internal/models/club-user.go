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
