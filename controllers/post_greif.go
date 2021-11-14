package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rwboyer/ginapi/models"
)

func PostGrief() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
		email_template := "SELECT id, subject, email_content FROM tbl_email_template WHERE expire IS NULL ORDER BY ID;"
		rows, err = models.Db.Query(email_template)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for rows.Next() {
			err := rows.Scan(
				&et.Id,
				&et.Subject,
				&et.Content,
				//&et.Header,
				//&et.Title,
				//&et.FilePath,
				//&et.FileName,
			)
			ets = append(ets, et)
			//log.Printf("ROW: %v", et)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		defer rows.Close()

		var tasks []interface{}
		var placeholders []string
		tdt := time.Now()
		for i, et := range ets {
			placeholders = append(placeholders,
				fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)",
					i*5+1,
					i*5+2,
					i*5+3,
					i*5+4,
					i*5+5,
				),
			)
			tasks = append(tasks, gr.Name, gr.Email, tdt, int(id), et.Id)
			tdt = tdt.AddDate(0, 0, 7)
		}
		log.Printf("TASKS: %v", tasks)

		task_insert := fmt.Sprintf("INSERT INTO tbl_task ( user_name, email_id, task_date, user_id, template_id ) VALUES %s",
			strings.Join(placeholders, ","),
		)

		_, err = models.Db.Exec(task_insert, tasks...)
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(gr)
		w.Write(j)

	}
}
