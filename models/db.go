package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var Db = opendb(os.Getenv("APIDB"))

func opendb(dbstring string) (*sql.DB) {
	db, err := sql.Open("mysql", dbstring)
	if( err != nil){
		fmt.Print(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)	

	return db
}
