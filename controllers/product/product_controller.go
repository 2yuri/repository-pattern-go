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
	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.ShowAllProducts()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, result)
}

func Create(c *gin.Context) {
	var dto product.ProductDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.CreateProduct(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, result)
}

func Update(c *gin.Context) {
	var dto product.ProductDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	err = service.UpdateProduct(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.Status(204)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.DeleteProduct(uint(newid))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, result)
}
