package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	_ "log"
	"net/http"

	_ "github.com/go-chi/chi/v5"
	"github.com/rwboyer/ginapi/models"
)

func PostObitDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vigils []models.Vigil
		var vigil models.Vigil

		//obit := chi.URLParam(r, "ref")
		o := r.URL.Query().Get("ref")

		log.Println(o)
	
		rows, err := models.Db.Query("select * from vigil_log where obit = ?;", o)
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
