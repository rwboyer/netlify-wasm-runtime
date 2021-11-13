package main

import (
	"github.com/rwboyer/ginapi/mappings"
	"net/http"
)

func main() {
	//mappings.Router.LoadHTMLGlob("templates/*.tmpl")
	mappings.CreateUrlMappings()
	http.ListenAndServe(":1111", mappings.Router)
}
