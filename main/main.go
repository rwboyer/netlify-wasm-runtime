package main

import(
	"github.com/rwboyer/ginapi/mappings"
)

func main () {
	mappings.CreateUrlMappings()
	mappings.Router.Run(":1111")
}