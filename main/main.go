package main

import(
	"github.com/rwboyer/ginapi/mappings"
)

func main () {
	//mappings.Router.LoadHTMLGlob("templates/*.tmpl")
	mappings.CreateUrlMappings()
	mappings.Router.Run(":1111")
}