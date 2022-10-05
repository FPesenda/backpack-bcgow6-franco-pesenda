package handler

import (
	"net/http"
	"strconv"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_1/PUT/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
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

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (prod *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := prod.service.Store(req.Name, req.Color, req.Code, req.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (prod *Product) UpdateByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDACIÖN DEL TOKEN
		token := ctx.Request.Header.Get("token")
		if token == "12345" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "no estas autorizado a realizar esta acción",
			})
		}
		//OBTENCIÓN DEL ID DESDE UN PARAMETRO
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "ID invalido",
			})
		}
		//TOMO EL VALOR DEL CUERPO PARA ALMACENARLO EN UNA VARIABLE
		var request request

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		if request.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "el nombre es requerido",
			})
		}
		if request.Color == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "el color es requerido",
			})
		}
		if request.Code == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "el codigo es requerido",
			})
		}
		if request.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "el precio es requerdio",
			})
		}

		productTemporal, err := prod.service.UpdateByID(int(id), request.Name, request.Color, request.Code, request.Price)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, productTemporal)
	}
}
