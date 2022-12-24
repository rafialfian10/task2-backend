package models

type Transaction struct {
	Id         int          `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int          `json:"qty" form:"qty" gorm:"type: int"`
	Total      int          `json:"total" form:"total" gorm:"type: int"`
	Status     string       `json:"status" form:"status" gorm:"type: varchar(255)"`
	Attachment string       `json:"attachment" form:"attachment" gorm:"type: varchar(255)"`
	TripId     int          `json:"-"`
	Trip       TripResponse `json:"trip" gorm:"foreignKey:trip_id"`
}
