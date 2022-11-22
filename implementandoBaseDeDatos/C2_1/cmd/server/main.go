package main

import (
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/cmd/server/handler"
	cnn "github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/db"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")
	{
		pr.POST("/", p.Store())
		pr.GET("/", p.GetByName())
		pr.GET("", p.GetAll())
		pr.DELETE("/:id", handler.IdValidationMiddleWare(), p.Delete())
		pr.PUT("/:id", handler.IdValidationMiddleWare(), p.Update())
	}

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
