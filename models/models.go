package models

import (
	"time"
)

type Vigil struct {
	Id     int       `json:"cid"`
	Date   time.Time `json:"date"`
	Obit   string    `json:"obit"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`
	Text   string    `json:"text"`
	Candle string    `json:"candle"`
	Img    string    `json:"img"`
}
