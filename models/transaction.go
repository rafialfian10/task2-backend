package models

type Transaction struct {
	Id         int          `json:"id"`
	CounterQty int          `json:"qty" form:"qty" gorm:"type: int"`
	Total      int          `json:"total" form:"total" gorm:"type: int"`
	Status     string       `json:"status" form:"status" gorm:"type: varchar(255)"`
	Attacment  string       `json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	TripId     int          `json:"trip_id"`
	Trip       TripResponse `json:"trip"`
}
