package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rwboyer/ginapi/models"
)


func GetObit() gin.HandlerFunc {
	return func (c *gin.Context) {
	var vigils []models.Vigil
	var vigil models.Vigil

	rows, err := models.Db.Query("select * from vigil_log")
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
}

