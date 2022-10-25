package handler

import (
	"net/http"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoWeb/C2_2/GenerateInternalPackage/internal/products"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Id    int
	Stock int
	Name  string
	Color string
	Code  string
	Price float64
}

type Product struct {
	service products.Service
}

func NewProduct(serv products.Service) *Product {
	return &Product{
		service: serv,
	}
}

func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}
