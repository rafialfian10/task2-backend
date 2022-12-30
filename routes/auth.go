package routes

import (
	"project/handlers"
	"project/pkg/mysql"
	"project/repositories"

	"github.com/gorilla/mux"
)

// Create AuthRoutes function and add "/register" route with handler Register and POST method here ...
func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/register_admin", h.RegisterAdmin).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}
