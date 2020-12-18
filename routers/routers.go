package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nilerajput91/assig5bookapi/controllers"
)

// SetupRouter to all routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("book", controllers.ListBook)
		v1.POST("book", controllers.AddNewBook)
		v1.GET("book/:id", controllers.GetOneBook)
		v1.PUT("book/:id", controllers.PutOneBook)
		v1.DELETE("book/:id", controllers.DeleteBook)
	}
	return router
}
