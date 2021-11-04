package controllers

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nfnt/resize"
)

func ImgPost() gin.HandlerFunc {
	return func (c *gin.Context){
		file, err := c.FormFile("img")
		if err != nil {
			log.Fatal(err)
		}
		
		log.Println(file.Filename)
		dir, _ := os.Getwd()
		log.Println(dir)

		f, _ := file.Open()
		defer f.Close()
		imData, _, _ := image.Decode(f)
		newImage := resize.Resize(600, 0, imData, resize.Lanczos3)	

		tempFile, _ := ioutil.TempFile("saved", "upload-*.jpg")
		defer tempFile.Close()
		jpeg.Encode(tempFile, newImage, &jpeg.Options{Quality: 50})

		c.SaveUploadedFile(file, "saved/"+file.Filename)

		c.JSON(http.StatusOK, gin.H{
			"message":  "file uploaded successfully",
		})
	}
}
