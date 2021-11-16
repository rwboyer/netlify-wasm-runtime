package controllers

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	_ "mime/multipart"
	"net/http"
	_ "net/smtp"
	"os"
	"time"

	"github.com/go-chi/httplog"
	"github.com/nfnt/resize"
	"github.com/rwboyer/ginapi/models"
)

func PostObitDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vigil models.Vigil

		oplog := httplog.LogEntry(r.Context())

		o := r.URL.Query().Get("ref")

		r.ParseMultipartForm(32 << 20)

		vigil.Name = r.MultipartForm.Value["name"][0]
		vigil.Email = r.MultipartForm.Value["email"][0]
		vigil.Phone = r.MultipartForm.Value["phone"][0]
		if r.MultipartForm.Value["candle"] != nil {
			vigil.Candle = "on"
		} else {
			vigil.Candle = "off"
		}
		vigil.Text = r.MultipartForm.Value["message"][0]
		vigil.Obit = o

		fheader := r.MultipartForm.File["pic"]
		if fheader == nil {
			oplog.Info().Msg("no file uploaded")
		} else {

			oplog.Info().Msg(fheader[0].Filename)

			f, _ := fheader[0].Open()
			defer f.Close()
			imData, _, _ := image.Decode(f)
			newImage := resize.Resize(600, 0, imData, resize.Lanczos3)

			t := time.Now()
			d := fmt.Sprintf("saved/%04d%02d", t.Year(), t.Month())
			os.Mkdir(d, 0755)
			tempFile, _ := ioutil.TempFile(d, "upload-*.jpg")
			defer tempFile.Close()
			jpeg.Encode(tempFile, newImage, &jpeg.Options{Quality: 50})
			vigil.Img = tempFile.Name()

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
			oplog.Err(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		j, _ := json.Marshal(vigil)
		w.Write(j)
	}
}
