package product

import (
	"strconv"

	repository "github.com/hyperyuri/repository-pattern-go/repositories/postgres"
	"github.com/hyperyuri/repository-pattern-go/services/product"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.ShowProduct(uint(newid))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, result)
}

func ShowAll(c *gin.Context) {

}
