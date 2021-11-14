package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rwboyer/ginapi/util"
)

func PostPreplan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result map[string]interface{}

		defer r.Body.Close()
		b, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(b, &result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		log.Printf("Raw map: %v", result)

		templ, err := util.LoadPrePlanT()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var buf bytes.Buffer
		templ.Execute(&buf, &result)

		var hdrs = map[string]string{}
		to := make([]string, 0)
		to = append(to, "rwboyer@mac.com") //McCrery Address anne@mccreryandharra.com
		to = append(to, fmt.Sprintf("%v", result["email"]))

		tm, err := util.NewHtmlMailer(to, "preplanning@mccreryharra.com", "Testing HTML", "", &hdrs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if err = tm.Send(buf.String()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(result)
		w.Write(j)
	}
}
