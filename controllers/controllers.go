package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"io/ioutil"
	"time"
	"log"
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

func GetObit(c *gin.Context) {
	var vigils []models.Vigil
	var vigil models.Vigil

	rows, err := db.Query("select * from vigil_log")
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
			if( err != nil){
				fmt.Println(err.Error())
			}
			vigils = append(vigils, vigil)
	}
	defer rows.Close()
	c.JSON(http.StatusOK, vigils)
}

func GetObitDetail(c *gin.Context){
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

func ImgPost(c *gin.Context){
	file, err := c.FormFile("img")
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(file.Filename)
	dir, _ := os.Getwd()
	log.Println(dir)

	f, _ := file.Open()
	defer f.Close()
	fileBytes, err := ioutil.ReadAll(f)
	tempFile, err := ioutil.TempFile("../saved", "upload-*.jpg")
	defer tempFile.Close()
	tempFile.Write(fileBytes)

	//log.Println(fileBytes)
	err = c.SaveUploadedFile(file, "../saved/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		//"pathname": u.EscapedPath(),
	})
}