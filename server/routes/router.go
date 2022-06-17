package routes

import (
	"firstgo-p/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controller.ShowAllBooks)
			books.GET("/:id", controller.ShowBook)
			books.POST("/", controller.CreateBook)
			books.PUT("/", controller.EditBook)
			books.DELETE("/:id", controller.DeleteBook)
		}
	}
	return router
}
