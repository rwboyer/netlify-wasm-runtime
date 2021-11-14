package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/rwboyer/ginapi/util"
)

func PostPreplan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result map[string]interface{}

		defer r.Body.Close()
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &result)

		log.Println(reflect.TypeOf(result))
		log.Println(result)

		templ, err := util.LoadPrePlanT()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var buf bytes.Buffer
		templ.Execute(&buf, result)

		var hdrs = map[string]string{}
		to := make([]string, 0)
		to = append(to, "rwboyer@mac.com")

		tm, err := util.NewHtmlMailer(to, "preplanning@mccreryharra.com", "Testing HTML", "", &hdrs)
		if err != nil {
			log.Println(err)
			return
		}

		if err = tm.Send(buf.String()); err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(result)
		w.Write(j)
	}
}
