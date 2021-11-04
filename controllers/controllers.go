package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db = opendb(os.Getenv("APIDB"))

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


func Cors() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

