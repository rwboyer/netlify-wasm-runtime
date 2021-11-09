package mappings

import (
	_ "net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rwboyer/ginapi/controllers"
)

var Router = chi.NewRouter()

func CreateUrlMappings() {
	//Router = gin.Default()
	//Router.LoadHTMLGlob("templates/*.tmpl")

	//Router.Use(controllers.Cors())
	Router.Use(middleware.Logger)
  Router.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))
	//v1 := Router.Group("/v1")
	Router.Get("/api/vigil/{id}*", controllers.GetObitDetail())
	Router.Post("/api/vigil", controllers.PostObitDetail())
	Router.Get("/api/vigil", controllers.GetObit())
	Router.Get("/hello/{name}", controllers.GetHello())
	//Router.Post("/img/", controllers.ImgPost())
	//Router.Post("/imgfun/", controllers.ImgPostFun())
	//Router.Static("/file", "saved")
}
