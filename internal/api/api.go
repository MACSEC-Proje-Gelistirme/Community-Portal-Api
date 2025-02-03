package api

import (
	"database/sql"

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

	//protected := router.PathPrefix("").Subrouter()
	//	protected.Use(middleware.EnsureValidToken)

	return router
}
