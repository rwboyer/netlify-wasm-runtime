package mappings

import(
	"github.com/rwboyer/ginapi/controllers" 
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine 

func CreateUrlMappings()  {
	Router = gin.Default()
	Router.LoadHTMLGlob("../templates/*.tmpl")


	Router.Use(controllers.Cors())
	v1 := Router.Group("/v1")
	{
		v1.GET("/obit/:id/*obit", controllers.GetObitDetail)
		v1.GET("/obit/", controllers.GetObit)
		v1.POST("/img/", controllers.ImgPost)
		v1.POST("/imgfun/", controllers.ImgPostFun)
	}
	Router.Static("/file", "../saved")
}