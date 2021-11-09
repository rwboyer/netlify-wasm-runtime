package controllers

import (
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
	_ "mime/multipart"
	_ "github.com/go-chi/chi/v5"
	"github.com/rwboyer/ginapi/models"
)

func PostObitDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vigil models.Vigil

		//obit := chi.URLParam(r, "ref")
		o := r.URL.Query().Get("ref")
		//json.NewDecoder(r.Body).Decode(&vigil)
		r.ParseMultipartForm(32 << 20)
		vigil.Name = r.MultipartForm.Value["name"][0]
		vigil.Email = r.MultipartForm.Value["email"][0]
		vigil.Phone = r.MultipartForm.Value["phone"][0]
		if r.MultipartForm.Value["candle"] != nil {
			vigil.Candle = "ON"
		} else{
			vigil.Candle = "OFF"
		}
		vigil.Text = r.MultipartForm.Value["message"][0]
		vigil.Obit = o

		
		log.Println(o)
	
		sqlStatement := `
		INSERT INTO vigil_log (obit, name, email, phone, text, candle, img)
		VALUES (?, ?, ?, ?, ?, ?, ?)`
		_, err := models.Db.Exec(sqlStatement, 
			vigil.Obit, 
			vigil.Name, 
			vigil.Email,
			vigil.Phone,
			vigil.Text,
			vigil.Candle,
			vigil.Img,
		)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(vigil)
		w.Write(j)
	}
}
