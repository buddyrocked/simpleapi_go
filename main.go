package main

import (
	"net/http"
	"simpleapi_go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/product", getProducts)
	router.POST("/product", addProduct)
	router.GET("/product/:code", getProduct)
	router.PUT("/product/:code", updateProduct)
	router.DELETE("/product/:code", deleteProduct)

	router.Run("localhost:8083")
}

func getProducts(c *gin.Context) {
	products := models.GetProducts()

	if products == nil || len(products) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "List products is empty"})
	} else {
		c.IndentedJSON(http.StatusOK, products)
	}
}

func getProduct(c *gin.Context) {
	code := c.Param("code")

	product := models.GetProduct(code)

	if product == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product " + code + " not found"})
	} else {
		c.IndentedJSON(http.StatusOK, product)
	}
}

func addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}

}

func updateProduct(c *gin.Context) {
	code := c.Param("code")
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product " + code + " not found"})
	} else {
		models.UpdateProduct(code, prod)
		c.IndentedJSON(http.StatusOK, prod)
	}

}

func deleteProduct(c *gin.Context) {
	code := c.Param("code")
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product " + code + " not found"})
	} else {
		models.DeleteProduct(code)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Product " + code + " deleted"})
	}

}
