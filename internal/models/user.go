package models

type User struct {
	UserID               string `json:"id"`
	SchoolId             string `json:"school_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Email                string `json:"email"`
	TelephoneNumber      string `json:"telephone_number"`
	EmailPreferences     bool   `json:"email_preferences"`
	MarketingPreferences bool   `json:"marketing_preferences"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
}

type CreateUserPayload struct {
	UserID          string `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	SchoolNumber    string `json:"school_number"`
	TelephoneNumber string `json:"telephone_number"`
	Email           string `json:"email"`
}
