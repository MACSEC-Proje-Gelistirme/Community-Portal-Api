package api

import (
	"database/sql"
	"net/http"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/middleware"
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

	router.Use(middleware.CorsMiddleware)

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.EnsureValidToken)

	protected.HandleFunc("/user", r.CreateUser).Methods(http.MethodPost, http.MethodOptions)

	return router
}
