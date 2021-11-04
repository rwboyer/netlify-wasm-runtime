package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rwboyer/ginapi/models"
	"net/http"
)

func GetObitDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		var vigils []models.Vigil
		var vigil models.Vigil

		obit := c.Param("obit")
		rows, err := models.Db.Query("select * from vigil_log where obit = ?;", obit)
		if err != nil {
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
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, vigils)
	}
}
