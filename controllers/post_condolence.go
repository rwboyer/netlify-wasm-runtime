package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rwboyer/ginapi/models"
	"github.com/rwboyer/ginapi/util"
)

func PostCondolence() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Data struct {
			Data models.Condolence `json:"data"`
		}

		oplog := httplog.LogEntry(r.Context())

		data := Data{}
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {

			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode("boom")
			oplog.Err(err)
			return
		}

		err = util.CheckRecaptcha(data.Data.Gresp)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode("Recaptcha Fail")
			oplog.Err(err)
			return
		}

		header := make(map[string]string)
		header["reply-to"] = data.Data.Email

		message := fmt.Sprintf("Message: %s\n\nPhone: %s\n", data.Data.Message, data.Data.Phone)

		from := "condolences@mccreryandharra.com"
		to := make([]string, 0)
		to = append(to, "rwboyer@mac.com") //data.Data.To)

		tm, err := util.NewTextMailer(to, from, fmt.Sprintf("Condolence from %s", data.Data.Name), message, &header)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = tm.Send(message); err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sqlStatement := `
		INSERT INTO condolence_log (obit, name, email, phone, text)
		VALUES (?, ?, ?, ?, ?)`
		_, err = models.Db.Exec(sqlStatement,
			data.Data.Obit,
			data.Data.Name,
			data.Data.Email,
			data.Data.Phone,
			data.Data.Message,
		)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(data.Data)
		w.Write(j)
	}
}
