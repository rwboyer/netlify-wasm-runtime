package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rwboyer/ginapi/models"
)

func GetObit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vigils []models.Vigil
		var vigil models.Vigil

		rows, err := models.Db.Query("select * from vigil_log")
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
			if err != nil {
				fmt.Println(err.Error())
			}
			vigils = append(vigils, vigil)
		}

		defer rows.Close()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(vigils)
		w.Write(j)
	}
}
