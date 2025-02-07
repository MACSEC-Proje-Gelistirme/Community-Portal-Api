package middleware

import (
	"database/sql"
	"net/http"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/permissions"
	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/repository"
	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/pkg/utils"
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

func CheckPermission(authService *AuthorizationService, permission permissions.Permission) func(http.Handler) http.Handler {
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
