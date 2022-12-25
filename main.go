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

	// route
	route := mux.NewRouter()

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// // route untuk menginisialisasi folder dengan file, image css, js agar dapat diakses kedalam project
	route.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// pathPrefix untuk membuat route baru. Subrouter untuk menguji route pada pathPrefix. RouteInit dari (routes/routes)
	routes.RouteInit(route.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:3000")
	http.ListenAndServe("localhost:3000", route)
}

// lifecycle: models ---> koneksi mysql ---> database migration ---> repositories ---> dto ---> handlers ---> routers
// run xampp: sudo /opt/lampp/lampp start
