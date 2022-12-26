package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project/database"
	"project/pkg/mysql"
	"project/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	var port = os.Getenv("PORT")

	fmt.Println("server running localhost:" + port)
	http.ListenAndServe("localhost:"+port, route)
}

// lifecycle: models ---> koneksi mysql ---> database migration ---> repositories ---> dto ---> handlers ---> routers
// run xampp: sudo /opt/lampp/lampp start
