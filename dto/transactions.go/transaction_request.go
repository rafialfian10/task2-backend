package transactionsdto

type CreateTransactionRequest struct {
	CounterQty int    `json:"qty" form:"qty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
	Image      string `json:"image" form:"image" validate:"required"`
	TripId     int    `json:"trip_id" form:"trip_id" validate:"required"`
}

type UpdateTransactionRequest struct {
	CounterQty int    `json:"qty" form:"qty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Image      string `json:"image" form:"image"`
	TripId     int    `json:"trip_id" form:"trip_id" validate:"required"`
}
