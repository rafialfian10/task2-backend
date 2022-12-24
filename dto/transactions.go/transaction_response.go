package transactionsdto

type TransactionResponse struct {
	Id         int    `json:"id"`
	CounterQty int    `json:"qty" form:"qty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attacment  string `json:"attachment" form:"attachment"`
	TripId     int    `json:"trip_id" form:"trip_id"`
}
