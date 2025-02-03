package middleware

import (
	"database/sql"
	"net/http"

	"github.com/btk-hackathon-24-debug-duo/project-setup/internal/permissions"
	"github.com/btk-hackathon-24-debug-duo/project-setup/internal/repository"
	"github.com/btk-hackathon-24-debug-duo/project-setup/pkg/utils"
)

type AuthorizationService struct {
	db *sql.DB
}

func newAuthorizationService(db *sql.DB) *AuthorizationService {
	return &AuthorizationService{db: db}
}

func (a *AuthorizationService) GetUserRole(clubID, userID string) (*permissions.Role, error) {
	clubRolesRepository := repository.NewClubUserRolesRepository(a.db)
	roleName, err := clubRolesRepository.GetUserRole(clubID, userID)
	if err != nil {
		return nil, err
	}

	return permissions.GetRoleWithRoleName(roleName), nil
}

func (a *AuthorizationService) HasPermission(role *permissions.Role, permission permissions.Permission) bool {
	return role.HasPermission(permission)
}

func checkPermission(authService *AuthorizationService, permission permissions.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := utils.GetTokenClaims(r)
			if !ok {
				utils.JSONError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			userID, ok := utils.GetUserIDFromClaims(claims)
			if !ok {
				utils.JSONError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			clubID := r.Header.Get("club-id")
			if clubID == "" {
				utils.JSONError(w, http.StatusBadRequest, "Club ID is required")
				return
			}

			role, err := authService.GetUserRole(clubID, userID)
			if err != nil || role == nil {
				utils.JSONError(w, http.StatusInternalServerError, "Unable to get user role")
				return
			}

			if !authService.HasPermission(role, permission) {
				utils.JSONError(w, http.StatusForbidden, "Forbidden")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
