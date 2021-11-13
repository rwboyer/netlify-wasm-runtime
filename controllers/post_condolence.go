package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-chi/chi/v5"
	"github.com/rwboyer/ginapi/models"
	"github.com/rwboyer/ginapi/util"
	"log"
	"net/http"
	"net/smtp"
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
		}

		header := make(map[string]string)
		header["reply-to"] = data.Data.Email

		message := fmt.Sprintf("Message: %s\n\nPhone: %s\n", data.Data.Message, data.Data.Phone)

		from := "condolences@mccreryandharra.com"
		to := make([]string, 0)
		to = append(to, "rwboyer@mac.com") //data.Data.To)

		smtp.SendMail("localhost", nil, from, to, []byte(message))

		util.Mailer(to, from, fmt.Sprintf("Condolence from %s", data.Data.Name), []byte(message), &header)

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
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(data.Data)
		w.Write(j)
	}
}
