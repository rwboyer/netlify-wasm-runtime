package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rwboyer/ginapi/util"
)

func ImgPostFun() gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := c.FormFile("img")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		f, _ := file.Open()
		defer f.Close()
		ascii_art, err := util.AsciiArt(f)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.HTML(http.StatusOK, "img.tmpl", gin.H{"Art": ascii_art})
	}
}
