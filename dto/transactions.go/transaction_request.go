package transactionsdto

type CreateTransactionRequest struct {
	CounterQty int    `json:"qty" form:"qty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
	Attachment string `json:"attachment" form:"attachment"`
	TripId     int    `json:"trip_id" form:"trip_id" validate:"required"`
}

type UpdateTransactionRequest struct {
	CounterQty int    `json:"qty" form:"qty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	TripId     int    `json:"trip_id" form:"trip_id" validate:"required"`
}
