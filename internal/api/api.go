package api

import (
	"database/sql"
	"net/http"

	"api/internal/middleware"
	"api/internal/permissions"

	"github.com/gorilla/mux"
)

type Router struct {
	db *sql.DB
}

func NewRouter(db *sql.DB) *Router {
	return &Router{
		db: db,
	}
}

func (r *Router) NewRouter() *mux.Router {
	router := mux.NewRouter()
	authService := middleware.NewAuthorizationService(r.db)

	router.Use(middleware.CorsMiddleware)

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.EnsureValidToken)

	protected.HandleFunc("/user", r.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/club-user", r.GetClubWithUserID).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/club-user", middleware.CheckPermission(authService, permissions.AddClubUser)(r.AddClubUser)).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/club-user", middleware.CheckPermission(authService, permissions.DeleteClubUser)(r.RemoveClubUser)).Methods(http.MethodDelete, http.MethodOptions)
	protected.HandleFunc("/club-user", middleware.CheckPermission(authService, permissions.UpdateClubUser)(r.UpdateClubUserRole)).Methods(http.MethodPut, http.MethodOptions)

	// Club endpoints
	protected.HandleFunc("/club", r.CreateClub).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/clubs", r.ListClubs).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/club", r.GetClub).Methods(http.MethodGet, http.MethodOptions)
	protected.HandleFunc("/club", middleware.CheckPermission(authService, permissions.ClubUpdatePermission)(r.UpdateClub)).Methods(http.MethodPut, http.MethodOptions)
	protected.HandleFunc("/club", middleware.CheckPermission(authService, permissions.ClubDeletePermission)(r.DeleteClub)).Methods(http.MethodDelete, http.MethodOptions)

	return router
}
