package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func validate(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "1234" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar la petición solicitada"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "petición realizada"})
	return
}

func main() {
	router := gin.Default()
	productsRouter := router.Group("/products")
	{
		productsRouter.POST("/", validate)
	}
	router.Run()
}
