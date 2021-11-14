package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
)

func PostPreplan() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result map[string]interface{}

		defer r.Body.Close()
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &result)

		log.Println(reflect.TypeOf(result))
		log.Println(result)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(result)
		w.Write(j)
	}
}
