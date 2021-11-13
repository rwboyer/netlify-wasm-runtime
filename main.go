package main

import (
	"net/http"
	"github.com/rwboyer/ginapi/mappings"
)

func main() {
	//mappings.Router.LoadHTMLGlob("templates/*.tmpl")
	mappings.CreateUrlMappings()
	http.ListenAndServe(":1111", mappings.Router)
}
