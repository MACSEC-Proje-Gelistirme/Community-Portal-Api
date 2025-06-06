package api

import (
	"net/http"

	"api/internal/models"
	"api/internal/permissions"
	"api/internal/repository"
	"api/pkg/utils"
)

func (ro *Router) CreateClub(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateClubPayload
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

	club := models.Club{
		Name:        payload.Name,
		Description: payload.Description,
		Email:       payload.Email,
	}

	clubRepository := repository.NewClubRepository(ro.db)
	clubID, err := clubRepository.CreateClub(club)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create club role for the creator as admin
	clubUserRepository := repository.NewClubUserRepository(ro.db)
	_, err = clubUserRepository.CreateClubRole(clubID, userID, "owner")
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{"id": clubID})
}

func (ro *Router) GetClub(w http.ResponseWriter, r *http.Request) {
	clubID := r.Header.Get("club-id")

	clubRepository := repository.NewClubRepository(ro.db)
	club, err := clubRepository.GetClubByID(clubID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if club == nil {
		utils.JSONError(w, http.StatusNotFound, "club not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, club)
}

func (ro *Router) UpdateClub(w http.ResponseWriter, r *http.Request) {
	clubID := r.Header.Get("club-id")

	var payload models.CreateClubPayload
	if err := utils.DecodeRequestBody(r, &payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	club := models.Club{
		ID:          clubID,
		Name:        payload.Name,
		Description: payload.Description,
		Email:       payload.Email,
	}

	clubRepository := repository.NewClubRepository(ro.db)
	err := clubRepository.UpdateClub(club)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "club updated successfully"})
}

func (ro *Router) DeleteClub(w http.ResponseWriter, r *http.Request) {
	clubID := r.Header.Get("club-id")

	clubUserRepository := repository.NewClubUserRepository(ro.db)

	userRole := r.Context().Value("userRole")

	if userRole != permissions.OwnerRole.Name {
		utils.JSONError(w, http.StatusBadRequest, "Only owner can delete club")
		return
	}

	err := clubUserRepository.DeleteAllClubRoles(clubID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	clubRepository := repository.NewClubRepository(ro.db)
	err = clubRepository.DeleteClub(clubID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "club deleted successfully"})
}

func (ro *Router) ListClubs(w http.ResponseWriter, r *http.Request) {
	clubRepository := repository.NewClubRepository(ro.db)
	clubs, err := clubRepository.ListClubs()
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, clubs)
}
