package repository

import (
	"api/internal/models"
	"database/sql"
	"time"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (e *EventRepository) CreateEvent(event *models.Event) (*models.Event, error) {
	var newEvent models.Event
	err := e.db.QueryRow(`
		INSERT INTO events (club_id, title, description, start_date, end_date, tags, location, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $8)
		RETURNING id, club_id, title, description, start_date, end_date,tags, location, created_at, updated_at`,
		event.ClubID, event.Title, event.Description, event.StartDate, event.EndDate, event.Tags, event.Location, time.Now(),
	).Scan(
		&newEvent.ID,
		&newEvent.ClubID,
		&newEvent.Title,
		&newEvent.Description,
		&newEvent.StartDate,
		&newEvent.EndDate,
		&newEvent.Tags,
		&newEvent.Location,
		&newEvent.CreatedAt,
		&newEvent.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &newEvent, nil
}

func (e *EventRepository) GetEventByID(eventID string) (*models.Event, error) {
	var event models.Event
	err := e.db.QueryRow(`
		SELECT id, club_id, title, description, start_date, end_date, tags, location, created_at, updated_at
		FROM events
		WHERE id = $1`, eventID,
	).Scan(
		&event.ID,
		&event.ClubID,
		&event.Title,
		&event.Description,
		&event.StartDate,
		&event.EndDate,
		&event.Tags,
		&event.Location,
		&event.CreatedAt,
		&event.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *EventRepository) GetAllEvents() ([]models.Event, error) {
	rows, err := e.db.Query(`
		SELECT id, club_id, title, description, start_date, end_date, tags, location, created_at, updated_at
		FROM events
		ORDER BY start_date DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.ID,
			&event.ClubID,
			&event.Title,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.Tags,
			&event.Location,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (e *EventRepository) UpdateEvent(eventID string, event *models.UpdateEventPayload) (*models.Event, error) {
	var updatedEvent models.Event
	err := e.db.QueryRow(`
		UPDATE events 
		SET title = $1, description = $2, start_date = $3, end_date = $4, location = $5, updated_at = $6
		WHERE id = $7
		RETURNING id, club_id, title, description, start_date, end_date, location, created_at, updated_at`,
		event.Title, event.Description, event.StartDate, event.EndDate, event.Location, time.Now(), eventID,
	).Scan(
		&updatedEvent.ID,
		&updatedEvent.ClubID,
		&updatedEvent.Title,
		&updatedEvent.Description,
		&updatedEvent.StartDate,
		&updatedEvent.EndDate,
		&updatedEvent.Location,
		&updatedEvent.CreatedAt,
		&updatedEvent.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &updatedEvent, nil
}

func (e *EventRepository) DeleteEvent(eventID string) error {
	_, err := e.db.Exec(`DELETE FROM events WHERE id = $1`, eventID)
	if err != nil {
		return err
	}
	return nil
}
