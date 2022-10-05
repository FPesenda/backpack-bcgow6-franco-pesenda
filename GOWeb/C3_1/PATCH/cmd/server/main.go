package main

import (
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_1/PATCH/cmd/server/handler"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_1/PATCH/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	//CRUD de PRODUCTOS
	repository := products.NewRepository()
	service := products.NewService(repository)
	products := handler.NewProduct(service)
	//CREACIÓN DEL ROUTER PRINCIPAL
	router := gin.Default()
	//CREAMOS EN EL RP UN GRUPO DE CONSULTAS
	productsRouter := router.Group("/products")
	{
		productsRouter.DELETE("/:id", products.Delete())
		productsRouter.GET("/", products.GetAll())
		productsRouter.POST("/", products.Store())
		productsRouter.PUT("/:id", products.UpdateByID())
	}
	//ARRANCAMOS EL ROUTER
	router.Run()
}
