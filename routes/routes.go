package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	TripRoutes(r)
	AuthRoutes(r)
	CountryRoutes(r)
}
