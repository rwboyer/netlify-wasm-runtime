package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rwboyer/ginapi/models"
)

func GetObitDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vigils []models.Vigil
		var vigil models.Vigil

		oplog := httplog.LogEntry(r.Context())
		obit := chi.URLParam(r, "*")

		rows, err := models.Db.Query("select * from vigil_log where obit = ?;", "/"+obit)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

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
				oplog.Err(err)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(vigils)
		w.Write(j)
	}
}
