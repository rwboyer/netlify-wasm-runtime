package controllers

import (
	"encoding/json"
	_ "fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	_ "mime/multipart"
	"net/http"
	_ "net/smtp"
	"os"

	_ "github.com/go-chi/chi/v5"
	"github.com/nfnt/resize"
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
			vigil.Candle = "on"
		} else{
			vigil.Candle = "off"
		}
		vigil.Text = r.MultipartForm.Value["message"][0]
		vigil.Obit = o

		log.Println(o)

		fheader := r.MultipartForm.File["pic"]
		if fheader == nil {
			log.Println("no upload")
		} else {

		log.Println(fheader[0].Filename)
		dir, _ := os.Getwd()
		log.Println(dir)

		f, _ := fheader[0].Open()
		defer f.Close()
		imData, _, _ := image.Decode(f)
		newImage := resize.Resize(600, 0, imData, resize.Lanczos3)

		tempFile, _ := ioutil.TempFile("saved", "upload-*.jpg")
		defer tempFile.Close()
		jpeg.Encode(tempFile, newImage, &jpeg.Options{Quality: 50})
		vigil.Img = tempFile.Name()

		//r.SaveUploadedFile(file, "saved/"+file.Filename)
		}

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
