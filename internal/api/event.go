package api

import (
	"api/internal/models"
	"api/internal/repository"
	"api/pkg/utils"
	"encoding/json"
	"net/http"
)

func (ro *Router) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateEventPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	clubID := r.Header.Get("club-id")

	event := models.Event{
		ClubID:      clubID,
		Title:       payload.Title,
		Description: payload.Description,
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		Tags:        payload.Tags,
		Location:    payload.Location,
	}

	eventRepository := repository.NewEventRepository(ro.db)
	newEvent, err := eventRepository.CreateEvent(&event)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, newEvent)
}

func (ro *Router) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.Header.Get("club-id")
	if eventID == "" {
		utils.JSONError(w, http.StatusBadRequest, "event id is required")
		return
	}

	eventRepository := repository.NewEventRepository(ro.db)
	event, err := eventRepository.GetEventByID(eventID)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "event not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, event)
}

func (ro *Router) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	eventRepository := repository.NewEventRepository(ro.db)
	events, err := eventRepository.GetAllEvents()
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, events)
}

func (ro *Router) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.Header.Get("event-id")
	if eventID == "" {
		utils.JSONError(w, http.StatusBadRequest, "event id is required")
		return
	}

	var event models.UpdateEventPayload
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	eventRepository := repository.NewEventRepository(ro.db)

	updatedEvent, err := eventRepository.UpdateEvent(eventID, &event)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, updatedEvent)
}

func (ro *Router) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := r.Header.Get("event-id")
	if eventID == "" {
		utils.JSONError(w, http.StatusBadRequest, "event id is required")
		return
	}

	eventRepository := repository.NewEventRepository(ro.db)
	_, err := eventRepository.GetEventByID(eventID)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "event not found")
		return
	}

	err = eventRepository.DeleteEvent(eventID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]bool{"success": true})
}
