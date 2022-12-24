package routes

import (
	"project/handlers"
	"project/pkg/middleware"
	"project/pkg/mysql"
	"project/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.UploadFile(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.UploadFile(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
}
