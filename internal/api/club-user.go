package api

import (
	"api/internal/models"
	"api/internal/permissions"
	"api/internal/repository"
	"api/pkg/utils"
	"encoding/json"
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

func (ro *Router) AddClubUser(w http.ResponseWriter, r *http.Request) {
	clubID := r.Header.Get("club-id")
	if clubID == "" {
		utils.JSONError(w, http.StatusBadRequest, "club id not found in header")
		return
	}

	var payload models.AddClubUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if !(payload.Role == permissions.AdminRole.Name ||
		payload.Role == permissions.ClubAdminRole.Name ||
		payload.Role == permissions.MailAdminRole.Name ||
		payload.Role == permissions.SocialAdminRole.Name) {
		utils.JSONError(w, http.StatusBadRequest, "invalid role")
		return
	}

	clubUserRepository := repository.NewClubUserRepository(ro.db)
	userRepository := repository.NewUserRepository(ro.db)

	// Get user by email
	user, err := userRepository.GetUserByEmail(payload.Email)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "user not found")
		return
	}

	success, err := clubUserRepository.CreateClubRole(clubID, user.UserID, payload.Role)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]bool{"success": success})
}

func (ro *Router) RemoveClubUser(w http.ResponseWriter, r *http.Request) {
	clubID := r.Header.Get("club-id")
	if clubID == "" {
		utils.JSONError(w, http.StatusBadRequest, "club id not found in header")
		return
	}

	var payload models.DeleteClubUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	clubUserRepository := repository.NewClubUserRepository(ro.db)
	userRepository := repository.NewUserRepository(ro.db)

	// Get user by email
	user, err := userRepository.GetUserByID(payload.UserID)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "user not found")
		return
	}

	userId := r.Context().Value("userId")
	role := r.Context().Value("userRole")

	if user.UserID == userId && role == permissions.OwnerRole.Name {
		utils.JSONError(w, http.StatusInternalServerError, "Owner cannot delete himself")
		return
	}

	// Remove user from club
	err = clubUserRepository.DeleteClubRole(clubID, user.UserID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]bool{"success": true})
}

func (ro *Router) UpdateClubUserRole(w http.ResponseWriter, r *http.Request) {
	clubID := r.Header.Get("club-id")
	if clubID == "" {
		utils.JSONError(w, http.StatusBadRequest, "club id not found in header")
		return
	}

	var payload models.UpdateClubUserRolePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if !(payload.Role == permissions.AdminRole.Name ||
		payload.Role == permissions.ClubAdminRole.Name ||
		payload.Role == permissions.MailAdminRole.Name ||
		payload.Role == permissions.SocialAdminRole.Name) {
		utils.JSONError(w, http.StatusBadRequest, "invalid role")
		return
	}

	clubUserRepository := repository.NewClubUserRepository(ro.db)
	userRepository := repository.NewUserRepository(ro.db)

	// Get user by ID
	user, err := userRepository.GetUserByID(payload.UserID)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "user not found")
		return
	}

	userId := r.Context().Value("userId")
	role := r.Context().Value("userRole")

	if user.UserID == userId && role == permissions.OwnerRole.Name {
		utils.JSONError(w, http.StatusInternalServerError, "Owner cannot update himself")
		return
	}

	// Update user role
	err = clubUserRepository.UpdateClubRole(clubID, user.UserID, payload.Role)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]bool{"success": true})
}

func (ro *Router) GetUserClubsWithRoles(w http.ResponseWriter, r *http.Request) {
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

	clubs, err := clubUserRepository.GetUserClubsWithRoles(userID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, clubs)
}

func (ro *Router) GetClubDetailsWithMembers(w http.ResponseWriter, r *http.Request) {
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

	clubID := r.Header.Get("club-id")
	if clubID == "" {
		utils.JSONError(w, http.StatusBadRequest, "club id is required")
		return
	}

	clubUserRepository := repository.NewClubUserRepository(ro.db)

	// Check if user has access to this club
	_, err := clubUserRepository.GetUserRole(clubID, userID)
	if err != nil {
		utils.JSONError(w, http.StatusForbidden, "access denied")
		return
	}

	// Get club details and members
	club, members, err := clubUserRepository.GetClubDetailsWithMembers(clubID)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := models.ClubDetailsResponse{
		Club:    *club,
		Members: members,
	}

	utils.JSONResponse(w, http.StatusOK, response)
}
