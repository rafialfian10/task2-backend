package tripsdto

import "time"

type TripResponse struct {
	Id             int       `json:"id"`
	Title          string    `json:"title" form:"title" validate:"required"`
	CountryId      int       `json:"country_id"`
	Accomodation   string    `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string    `json:"transportation" form:"transportation" validate:"required"`
	Eat            string    `json:"eat" form:"eat" validate:"required"`
	Day            int       `json:"day" form:"day" validate:"required"`
	Night          int       `json:"night" form:"night" validate:"required"`
	DateTrip       time.Time `json:"datetrip"`
	Price          int       `json:"price" form:"price" validate:"required"`
	Quota          int       `json:"quota" form:"quota" validate:"required"`
	Description    string    `json:"description" form:"description" validate:"required"`
	Image          string    `json:"image" form:"image" validate:"required"`
}
