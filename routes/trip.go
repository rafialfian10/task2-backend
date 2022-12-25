package routes

import (
	"project/handlers"
	"project/pkg/middleware"
	"project/pkg/mysql"
	"project/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	TripRepository := repositories.RepositoriyTrip(mysql.DB)
	h := handlers.HandlerTrip(TripRepository)

	r.HandleFunc("/trips", h.FindTrips).Methods("GET")
	r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip))).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTrip))).Methods("PATCH")
	r.HandleFunc("/trip/{id}", middleware.Auth(h.DeleteTrip)).Methods("DELETE")
}
