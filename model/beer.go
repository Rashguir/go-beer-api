package model

import "time"

// Beer ...
type Beer struct {
	PK        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	ID string `json:"id"`
	Name string `json:"name"`
	ABV float32 `json:"abv"`
	BreweryID string `json:"brewery_id"`
	BreweryName string `json:"name_breweries"`
	City string `json:"city"`
	Country string `json:"country"`
	State string `json:"state"`
}
