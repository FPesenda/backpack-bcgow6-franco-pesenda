package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	BBDD []products
)

type products struct {
	Id          int
	Stock       int
	Name        string  `binding:"required"`
	Color       string  `binding:"required"`
	Code        string  `binding:"required"`
	Price       float64 `binding:"required"`
	Publish     bool    `binding:"required"`
	DateCreatin string  `binding:"required"`
}

func addProduct(ctx *gin.Context) {

	var productIn products

	token := ctx.GetHeader("token")

	if token != "1234" {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "No autorizado"})
		return
	}

	if err := ctx.ShouldBindJSON(&productIn); err != nil {
		errs := validateFields(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}
	productIn.Id = len(BBDD) + 1
	BBDD = append(BBDD, productIn)
	ctx.JSON(http.StatusOK, gin.H{"status": "se pudo agregar el producto"})
	return

}

func validateFields(err error) (errs []string) {
	var validator validator.ValidationErrors
	if errors.As(err, &validator) {
		for _, fe := range validator {
			errs = append(errs, fmt.Sprint("El campo ", fe.Field(), " es requerido"))
		}
	}
	return
}

func main() {
	router := gin.Default()

	productsGroup := router.Group("/products")
	{
		productsGroup.POST("/", addProduct)
	}

	router.Run()
}
