package controllers

import (
	"bytes"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/rwboyer/ginapi/util"
)

func ImagePostFun(nameT string) http.HandlerFunc {

	asciiHtml, _ := util.LoadAsciiArtT(nameT)

	return func(w http.ResponseWriter, r *http.Request) {

		var ascii_art string
		var buf bytes.Buffer

		oplog := httplog.LogEntry(r.Context())

		r.ParseMultipartForm(32 << 20)
		fheader := r.MultipartForm.File["pic"]
		if fheader == nil {
			oplog.Error()
		} else {

			oplog.Info().Msg(fheader[0].Filename)

			f, _ := fheader[0].Open()
			defer f.Close()
			ascii_art, _ = util.AsciiArt(f)
			s := map[string]interface{}{"Art": ascii_art}
			asciiHtml.Execute(&buf, &s)
		}

		w.Write(buf.Bytes())
	}
}
