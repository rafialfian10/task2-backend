package models

type Country struct {
	Id   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name"`
}

// relasi dengan tabel lain
type CountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
