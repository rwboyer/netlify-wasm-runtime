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

type Preplan struct {
	LastName   string `json:"lastName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
}

type GriefUser struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Cdate  time.Time
	Pub    string
	Last   int
	Remove string
	Gresp  string `json:"gresponse"`
}

type GriefTemplate struct {
	Id       int
	Expire   time.Time
	Subject  string
	Content  string
	Header   string
	Title    string
	FilePath string
	FileName string
}

type GriefTask struct {
	Id         int
	UserName   string
	EmailId    string
	TaskDate   time.Time
	MailSent   string
	UserId     int
	TemplateId int
	SentDate   time.Time
}
