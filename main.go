package main

import (
	"net/http"
	"github.com/rwboyer/ginapi/mappings"
)

func main() {
	//mappings.Router.LoadHTMLGlob("templates/*.tmpl")
	mappings.CreateUrlMappings()
	//mappings.Router.Run(":1111")
	http.ListenAndServe(":1111", mappings.Router)
}
