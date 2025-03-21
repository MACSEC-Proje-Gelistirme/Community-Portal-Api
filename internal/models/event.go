package models

type Event struct {
	ID          string `json:"id"`
	ClubID      string `json:"club_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Tags        string `json:"tags"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateEventPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Tags        string `json:"tags"`
	Location    string `json:"location"`
}

type UpdateEventPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Tags        string `json:"tags"`
	Location    string `json:"location"`
}
