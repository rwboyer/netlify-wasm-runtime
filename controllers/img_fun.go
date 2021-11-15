package controllers

import (
	"bytes"
	"log"
	"net/http"

	"github.com/rwboyer/ginapi/util"
)

func ImagePostFun(nameT string) http.HandlerFunc {

	asciiHtml, _ := util.LoadAsciiArtT(nameT)

	return func(w http.ResponseWriter, r *http.Request) {

		var ascii_art string
		//var err error
		var buf bytes.Buffer

		r.ParseMultipartForm(32 << 20)
		fheader := r.MultipartForm.File["pic"]
		if fheader == nil {
			log.Println("no upload")
		} else {

			log.Println(fheader[0].Filename)

			f, _ := fheader[0].Open()
			defer f.Close()
			ascii_art, _ = util.AsciiArt(f)
			s := map[string]interface{}{"Art": ascii_art}
			asciiHtml.Execute(&buf, &s)
		}

		w.Write(buf.Bytes())
	}
}
