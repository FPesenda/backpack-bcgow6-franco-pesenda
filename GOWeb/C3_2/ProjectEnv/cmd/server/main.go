package main

import (
	"log"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_2/ProjectEnv/cmd/server/handler"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_2/ProjectEnv/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al querer cargar el archivo .env")
	}
	//CRUD de PRODUCTOS
	repository := products.NewRepository()
	service := products.NewService(repository)
	products := handler.NewProduct(service)
	//CREACIÓN DEL ROUTER PRINCIPAL
	router := gin.Default()
	//CREAMOS EN EL RP UN GRUPO DE CONSULTAS
	productsRouter := router.Group("/products")
	{
		productsRouter.PATCH("/:id", products.Patch())
		productsRouter.DELETE("/:id", products.Delete())
		productsRouter.GET("/", products.GetAll())
		productsRouter.POST("/", products.Store())
		productsRouter.PUT("/:id", products.UpdateByID())
	}
	//ARRANCAMOS EL ROUTER
	router.Run()
}
