package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/rwboyer/ginapi/models"
)

func GetCondolence() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cond models.Condolence
		var conds []models.Condolence

		oplog := httplog.LogEntry(r.Context())
		obit := chi.URLParam(r, "*")

		rows, err := models.Db.Query("select * from condolence_log where obit = ?;", "/"+obit)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(
				&cond.Id,
				&cond.Date,
				&cond.Obit,
				&cond.Name,
				&cond.Email,
				&cond.Phone,
				&cond.Message,
			)
			conds = append(conds, cond)
			if err != nil {
				oplog.Err(err)
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(conds)
		w.Write(j)
	}
}
