package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hyperyuri/repository-pattern-go/controllers/product"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.GET("/")
		}

		prod := main.Group("product")
		{
			prod.GET("/", product.ShowAll)
			prod.GET("/:id", product.Show)
		}
	}

	return router
}
