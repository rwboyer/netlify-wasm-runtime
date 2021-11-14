package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rwboyer/ginapi/models"
	"github.com/rwboyer/ginapi/util"
)

func PostCondolence() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Data struct {
			Data models.Condolence `json:"data"`
		}

		data := Data{}
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {

			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode("boom")
		}

		err = util.CheckRecaptcha(data.Data.Gresp)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode("Recaptcha Fail")
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
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if err = tm.Send(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
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
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(data.Data)
		w.Write(j)
	}
}
