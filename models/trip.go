package models

import "time"

type Trip struct {
	Id             int             `json:"id"`
	Title          string          `json:"title" form:"title" gorm:"type: varchar(255)"`
	CountryId      int             `json:"country_id" gorm:"type: int"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation" form:"accmodation" gorm:"type: varchar(255)"`
	Transportation string          `json:"transportation" form:"transportation" gorm:"type: varchar(255)"`
	Eat            string          `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Day            int             `json:"day" form:"day" gorm:"type: int"`
	Night          int             `json:"night" form:"night" gorm:"type: int"`
	DateTrip       time.Time       `json:"datetrip"`
	Price          int             `json:"price" form:"price" gorm:"type: int"`
	Quota          int             `json:"quota" form:"quota" gorm:"type: int"`
	Description    string          `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image          string          `json:"image" form:"image" gorm:"type: varchar(255)"`
}

// type ProductResponse struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Desc   string `json:"desc"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	UserID int    `json:"-"`
// 	// User       UsersProfileResponse `json:"user"`
// 	Country   []Country `json:"category" gorm:"hasOne:product_categories"`
// 	CountryID []int     `json:"-" form:"category_id" gorm:"-"`
// }

// type ProductUserResponse struct {
// 	ID     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Desc   string `json:"desc"`
// 	Price  int    `json:"price"`
// 	Image  string `json:"image"`
// 	Qty    int    `json:"qty"`
// 	UserID int    `json:"-"`
// }

// func (ProductResponse) TableName() string {
// 	return "products"
// }

// func (ProductUserResponse) TableName() string {
// 	return "products"
// }
