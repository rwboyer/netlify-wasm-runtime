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

type Condolence struct {
	Id      int       `json:"cid"`
	Date    time.Time `json:"date"`
	Obit    string    `json:"refer"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	To      string    `json:"to"`
	Phone   string    `json:"phone"`
	Message string    `json:"message"`
	Gresp   string    `json:"gresponse"`
}

type Preplan struct{
	LastName		string	`json:"lastName"`
	FirstName		string	`json:"firstName"`
	MiddleName	string	`json:"middleName"`
}
