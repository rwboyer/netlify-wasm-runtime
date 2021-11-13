package controllers

import(
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rwboyer/ginapi/models"
)

func GetCondolence() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cond models.Condolence
		var conds []models.Condolence

		obit := chi.URLParam(r, "*")

		rows, err := models.Db.Query("select * from condolence_log where obit = ?;", "/" + obit)
		if err != nil {
			fmt.Println(err.Error())
		}
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
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(conds)
		w.Write(j)
	}
}