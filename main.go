package main

import (
	"fmt"
	"net/http"
	"project/database"
	"project/pkg/mysql"
	"project/routes"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	routes.RouteInit(route.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:3000")
	http.ListenAndServe("localhost:3000", route)
}

// lifecycle: models ---> koneksi mysql ---> database migration ---> repositories ---> dto ---> handlers ---> routers
// run xampp: sudo /opt/lampp/lampp start
