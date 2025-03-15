package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"api/internal/permissions"
	"api/internal/repository"
	"api/pkg/utils"
)

type AuthorizationService struct {
	db *sql.DB
}

func NewAuthorizationService(db *sql.DB) *AuthorizationService {
	return &AuthorizationService{db: db}
}

func (a *AuthorizationService) GetUserRole(clubID, userID string) (*permissions.Role, error) {
	clubRolesRepository := repository.NewClubUserRepository(a.db)
	roleName, err := clubRolesRepository.GetUserRole(clubID, userID)
	if err != nil {
		return nil, err
	}

	return permissions.GetRoleWithRoleName(roleName), nil
}

func (a *AuthorizationService) HasPermission(role *permissions.Role, permission permissions.Permission) bool {
	return role.HasPermission(permission)
}

func CheckPermission(authService *AuthorizationService, permission permissions.Permission) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
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

			ctx := r.Context()
			ctx = context.WithValue(ctx, "userRole", role)
			ctx = context.WithValue(ctx, "userId", userID)
			ctx = context.WithValue(ctx, "clubId", clubID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
