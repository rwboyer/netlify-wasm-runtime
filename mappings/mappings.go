package mappings

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/rwboyer/ginapi/controllers"
)

var Router = chi.NewRouter()

func CreateUrlMappings() {
	//Router.LoadHTMLGlob("templates/*.tmpl")

	Router.Use(middleware.Logger)
  Router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "saved"))

	controllers.FileServer(Router, "/saved", filesDir)

	Router.Get("/vigil/{id}*", controllers.GetObitDetail())
	Router.Post("/vigil", controllers.PostObitDetail())
	Router.Get("/vigil", controllers.GetObit())
	Router.Get("/hello/{name}", controllers.GetHello())
	Router.Get("/condolence/{id}*", controllers.GetCondolence())
	Router.Post("/condolence", controllers.PostCondolence())
	Router.Post("/preplan", controllers.PostPreplan())
	//Router.Post("/img/", controllers.ImgPost())
	//Router.Post("/imgfun/", controllers.ImgPostFun())
	//Router.Static("/file", "saved")
}
