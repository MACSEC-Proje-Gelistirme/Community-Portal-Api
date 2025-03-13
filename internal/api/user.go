package api

import (
	"net/http"

	"api/internal/models"
	"api/internal/repository"
	"api/pkg/utils"
)

func (ro *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateUserPayload
	if err := utils.DecodeRequestBody(r, &payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, ok := utils.GetTokenClaims(r)
	if !ok {
		utils.JSONError(w, http.StatusBadRequest, "claims not found")
		return
	}

	userID, ok := utils.GetUserIDFromClaims(claims)
	if !ok {
		utils.JSONError(w, http.StatusBadRequest, "user id not found")
		return
	}

	user := models.User{
		UserID:               userID,
		FirstName:            payload.FirstName,
		LastName:             payload.LastName,
		Email:                payload.Email,
		TelephoneNumber:      payload.TelephoneNumber,
		EmailPreferences:     true,
		MarketingPreferences: true,
		CreatedAt:            utils.GetCurrentTime(),
		UpdatedAt:            utils.GetCurrentTime(),
	}

	userRepository := repository.NewUserRepository(ro.db)
	result, err := userRepository.CreateUser(user)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, result)
}
