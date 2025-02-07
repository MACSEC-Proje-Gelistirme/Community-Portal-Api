package api

import (
	"net/http"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/models"
	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/repository"
	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/pkg/utils"
)

func (ro *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := utils.GetTokenClaims(r)
	if !ok {
		utils.JSONError(w, http.StatusUnauthorized, "Token claims not found")
		return
	}

	userID, ok := utils.GetUserIDFromClaims(claims)
	if !ok {
		utils.JSONError(w, http.StatusUnauthorized, "User ID not found in token claims")
		return
	}

	var payload models.CreateUserPayload
	if err := utils.DecodeRequestBody(r, &payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error())
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
