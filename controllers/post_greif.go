package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/httplog"
	"github.com/rwboyer/ginapi/models"
)

func PostGrief() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		oplog := httplog.LogEntry(r.Context())

		gr := models.GriefUser{
			Cdate:  time.Now(),
			Pub:    "y",
			Last:   0,
			Remove: "n",
		}
		err := json.NewDecoder(r.Body).Decode(&gr)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id, _ := result.LastInsertId()
		oplog.Info().Msgf("GRIEF USER CREATED SQL ID: %05d", id)

		var rows *sql.Rows
		var et models.GriefTemplate
		var ets []models.GriefTemplate
		email_template := "SELECT id, subject, email_content FROM tbl_email_template WHERE expire IS NULL ORDER BY ID;"
		rows, err = models.Db.Query(email_template)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

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
				oplog.Err(err)
			}
		}

		var tasks []interface{}
		var placeholders []string
		tdt := time.Now()
		for _, et := range ets {
			placeholders = append(placeholders, "(?, ?, ?, ?, ?)")
			tasks = append(tasks, gr.Name, gr.Email, tdt, int(id), et.Id)
			tdt = tdt.AddDate(0, 0, 7)
		}

		task_insert := fmt.Sprintf("INSERT INTO tbl_task ( user_name, email_id, task_date, user_id, template_id ) VALUES %s",
			strings.Join(placeholders, ","),
		)

		_, err = models.Db.Exec(task_insert, tasks...)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(gr)
		w.Write(j)

	}
}
