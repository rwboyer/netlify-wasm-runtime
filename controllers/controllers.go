package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rwboyer/ginapi/models"
)

var db = opendb(os.Getenv("APIDB"))

func opendb(dbstring string) (*sql.DB) {
	db, err := sql.Open("mysql", dbstring)
	if( err != nil){
		fmt.Print(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)	
	//defer db.Close()

	return db
}


func Cors() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetUser(c *gin.Context) {
	var vigils []models.Vigil
	var vigil models.Vigil

	rows, err := db.Query("select cid, date, name, text from vigil_log;")
	if( err != nil){
		fmt.Println(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&vigil.Id, &vigil.Date, &vigil.Name, &vigil.Text)
		vigils = append(vigils, vigil)
		if( err != nil){
			fmt.Println(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, vigils)
}

func GetUserDetail(c *gin.Context){
	var vigils []models.Vigil
	var vigil models.Vigil

	obit := c.Param("obit")
	rows, err := db.Query("select * from vigil_log where obit = ?;", obit)
	if( err != nil){
		fmt.Println(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(
			&vigil.Id,
			&vigil.Date, 
			&vigil.Obit, 
			&vigil.Name,
			&vigil.Email,
			&vigil.Phone,
			&vigil.Text,
			&vigil.Candle,
			&vigil.Img)
		vigils = append(vigils, vigil)
		if( err != nil){
			fmt.Println(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, vigils)
}