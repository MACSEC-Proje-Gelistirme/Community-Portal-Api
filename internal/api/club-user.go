package api

import (
	"api/internal/repository"
	"api/pkg/utils"
	"net/http"
)

func (ro *Router) GetClubWithUserID(w http.ResponseWriter, r *http.Request) {
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

	clubUserRepository := repository.NewClubUserRepository(ro.db)

	response, err := clubUserRepository.GetClubsWithUserID(userID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, response)

}
