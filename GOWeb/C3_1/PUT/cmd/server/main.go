package main

import (
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_1/PUT/cmd/server/handler"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_1/PUT/internal/products"
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
		productsRouter.GET("/", products.GetAll())
		productsRouter.POST("/", products.Store())
		productsRouter.PUT("/:id", products.UpdateByID())
	}
	//ARRANCAMOS EL ROUTER
	router.Run()
}
