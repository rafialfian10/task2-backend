package models

import "time"

type User struct {
	Id        int             `json:"id"`
	Name      string          `json:"name" gorm:"type: varchar(255)"`
	Email     string          `json:"email" gorm:"type: varchar(255)"`
	Password  string          `json:"password" gorm:"type: varchar(255)"`
	Phone     string          `json:"phone" gorm:"type: varchar(255)"`
	Address   string          `json:"address" gorm:"type: varchar(255)"`
	Profile   ProfileResponse `json:"profile"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

// relasi dengan tabel lain
type UsersResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersResponse) TableName() string {
	return "users"
}
