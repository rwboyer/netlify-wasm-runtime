package controllers

import (
	"database/sql"
	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rwboyer/ginapi/models"
)

var db = opendb("root:root@root(127.0.0.1:8889)/mccrery_grief")

func opendb(dbstring string) (*sql.DB) {
	db, err := sql.Open("mysql", dbstring)
	if( err != nil){
		fmt.Print(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)	
	defer db.Close()

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

	newdb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/mccrery_grief?parseTime=true")
	if( err != nil){
		fmt.Println(err.Error())
	}
	rows, err := newdb.Query("select cid, date, name, text from vigil_log;")
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
	c.JSON(http.StatusOK, nil)
}