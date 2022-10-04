package routes

import (
	"github.com/JrSchmidtt/api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetAllBooks)
			books.POST("/", controllers.CreateBook)
			books.GET("/:id", controllers.GetBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}
