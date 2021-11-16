package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rwboyer/ginapi/util"
)

func PostPreplan(tmplName string) http.HandlerFunc {

	templ, err := util.LoadPrePlanT(tmplName)
	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var result map[string]interface{}

		oplog := httplog.LogEntry(r.Context())

		defer r.Body.Close()
		b, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(b, &result); err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		templ.Execute(&buf, &result)

		var hdrs = map[string]string{}
		to := make([]string, 0)
		to = append(to, "anne@mccreryandharra.com") //McCrery Address anne@mccreryandharra.com
		//to = append(to, "rwboyer@mac.com") //McCrery Address anne@mccreryandharra.com
		to = append(to, fmt.Sprintf("%v", result["email"]))

		tm, err := util.NewHtmlMailer(to,
			"preplanning@mccreryharra.com",
			"Testing HTML",
			"",
			&hdrs)
		if err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = tm.Send(buf.String()); err != nil {
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(result)
		w.Write(j)
	}
}
