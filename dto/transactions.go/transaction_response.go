package transactionsdto

type TransactionResponse struct {
	Id         int    `json:"id"`
	CounterQty int    `json:"qty" form:"qty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
	Attacment  string `json:"attachment" form:"attachment" validate:"required"`
}
