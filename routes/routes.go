package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hyperyuri/repository-pattern-go/controllers/product"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		prod := main.Group("product")
		{
			prod.GET("/", product.ShowAll)
			prod.GET("/:id", product.Show)
			prod.POST("/", product.Create)
			prod.PUT("/:", product.Update)
			prod.DELETE("/:id", product.Delete)
		}
	}

	return router
}
