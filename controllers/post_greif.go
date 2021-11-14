package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rwboyer/ginapi/models"
)

func PostGrief() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//task_insert := "INSERT INTO tbl_task ( user_name, email_id, task_date, user_id, template_id ) VALUES (:name, :email, :date, :id, :template)"

		gr := models.GriefUser{
			Cdate:  time.Now(),
			Pub:    "y",
			Last:   0,
			Remove: "n",
		}
		err := json.NewDecoder(r.Body).Decode(&gr)

		var result sql.Result
		name_insert := "INSERT INTO tbl_user (user_name, email_id, create_date, publish, last_template_sent, remove) VALUES (?, ?, ?, ?, ?, ?)"
		result, err = models.Db.Exec(name_insert,
			gr.Name,
			gr.Email,
			gr.Cdate,
			gr.Pub,
			gr.Last,
			gr.Remove,
		)
		if err != nil {
			log.Println(err)
			return
		}
		id, _ := result.LastInsertId()
		log.Printf("GRIEF USER CREATED SQL ID: %05d", id)

		var rows *sql.Rows
		var et models.GriefTemplate
		var ets []models.GriefTemplate
		email_template := "SELECT * FROM tbl_email_template WHERE expire IS NULL ORDER BY ID;"
		rows, err = models.Db.Query(email_template)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for rows.Next() {
			err := rows.Scan(
				&et.Id,
				&et.Expire,
				&et.Subject,
				&et.Content,
				&et.Header,
				&et.Title,
				&et.FilePath,
				&et.FileName,
			)
			ets = append(ets, et)
			log.Printf("ROW: %v", et)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(gr)
		w.Write(j)

	}
}
