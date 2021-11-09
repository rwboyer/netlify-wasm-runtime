package mappings

import (
	_ "net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rwboyer/ginapi/controllers"
)

var Router = chi.NewRouter()

func CreateUrlMappings() {
	//Router = gin.Default()
	//Router.LoadHTMLGlob("templates/*.tmpl")

	//Router.Use(controllers.Cors())
	Router.Use(middleware.Logger)
	//v1 := Router.Group("/v1")
	Router.Get("/obit/{id}*", controllers.GetObitDetail())
	Router.Get("/obit/", controllers.GetObit())
	Router.Get("/hello/{name}", controllers.GetHello())
	//Router.Post("/img/", controllers.ImgPost())
	//Router.Post("/imgfun/", controllers.ImgPostFun())
	//Router.Static("/file", "saved")
}
