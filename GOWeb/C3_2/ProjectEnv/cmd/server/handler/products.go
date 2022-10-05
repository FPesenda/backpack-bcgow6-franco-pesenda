package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C3_2/ProjectEnv/internal/products"
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

func tokenValidator(token string, ctx *gin.Context) {
	if token != os.Getenv("TOKEN") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}
}

func (prod *Product) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidator(ctx.Request.Header.Get("token"), ctx)
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		productTemporal, errPatch := prod.service.Patch(int(id), req.Name, req.Price)
		if errPatch != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": errPatch,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": productTemporal,
		})
	}
}
func (prod *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "ID invalido",
			})
			return
		}
		err = prod.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprint("El producto con id ", id, " no se encuentra en la BBDD"),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("El producto %d a sido eliminado correctamente", id),
		})
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
			return
		}

		if request.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "el nombre es requerido",
			})
			return
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
			return
		}
		if request.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "el precio es requerdio",
			})
			return
		}

		productTemporal, err := prod.service.UpdateByID(int(id), request.Name, request.Color, request.Code, request.Price)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, productTemporal)
	}
}
