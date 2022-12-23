package routes

import (
	"project/handlers"
	"project/pkg/mysql"
	"project/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	TripRepository := repositories.RepositoriyTrip(mysql.DB)
	h := handlers.HandlerTrip(TripRepository)

	r.HandleFunc("/trips", h.FindTrips).Methods("GET")
	r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trip", h.CreateTrip).Methods("POST")
	r.HandleFunc("/trip/{id}", h.UpdateTrip).Methods("PATCH")
	r.HandleFunc("/trip/{id}", h.DeleteTrip).Methods("DELETE")
}
