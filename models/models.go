package models

import (
	"time"
)

type Vigil struct {
	Id int
	Date time.Time
	Obit string
	Name string
	Email string
	Phone string
	Text string
	Candle string
	Img string
}