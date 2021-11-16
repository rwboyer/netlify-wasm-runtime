package controllers

import (
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, _ := json.Marshal(name)
		w.Write(j)
	}
}
