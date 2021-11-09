package controllers

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/rwboyer/ginapi/models"
)

func GetObitDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vigils []models.Vigil
		var vigil models.Vigil

		obit := chi.URLParam(r, "*")
	
		rows, err := models.Db.Query("select * from vigil_log where obit = ?;", "/" + obit)
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
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(vigils)
		w.Write(j)
	}
}
