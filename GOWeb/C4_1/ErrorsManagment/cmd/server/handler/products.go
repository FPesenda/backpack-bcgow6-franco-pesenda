package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C4_1/ErrorsManagment/internal/products"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GOWeb/C4_1/ErrorsManagment/pkg/web"
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
		ctx.JSON(http.StatusUnauthorized, web.NewResponse(
			http.StatusUnauthorized,
			nil,
			"Token invalido",
		))
		return
	}
}

func (prod *Product) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidator(ctx.Request.Header.Get("token"), ctx)
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				err.Error(),
			))
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"Error al intentar parsear a structure",
			))
			return
		}

		productTemporal, errPatch := prod.service.Patch(int(id), req.Name, req.Price)
		if errPatch != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"Erro del patch",
			))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(
			http.StatusOK,
			productTemporal,
			"",
		))
	}
}
func (prod *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenValidator(ctx.Request.Header.Get("token"), ctx)

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"ID invalido",
			))
			return
		}
		err = prod.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				fmt.Sprint("El producto con id ", id, " no se encuentra en la BBDD"),
			))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(
			http.StatusOK,
			fmt.Sprint("El producto con id ", id, " a sido eliminado correctamente"),
			"",
		))
	}
}

// GET ALL
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenValidator(ctx.Request.Header.Get("token"), ctx)

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				400,
				nil,
				"No hay elementos en la BBDD",
			))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(
			http.StatusOK,
			p,
			"",
		))
	}
}

func (prod *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenValidator(ctx.Request.Header.Get("token"), ctx)

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				err.Error(),
			))
			return
		}
		p, err := prod.service.Store(req.Name, req.Color, req.Code, req.Price)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				err.Error(),
			))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(
			http.StatusOK,
			p,
			"",
		))
	}
}

func (prod *Product) UpdateByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDACIÖN DEL TOKEN
		tokenValidator(ctx.Request.Header.Get("token"), ctx)
		//OBTENCIÓN DEL ID DESDE UN PARAMETRO
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"Id invalido",
			))
		}
		//TOMO EL VALOR DEL CUERPO PARA ALMACENARLO EN UNA VARIABLE
		var request request

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				err.Error(),
			))
			return
		}

		if request.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"El nombre es requerido",
			))
			return
		}
		if request.Color == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"El color es requerido",
			))
		}
		if request.Code == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"El codigo es requerido",
			))
			return
		}
		if request.Price == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(
				http.StatusBadRequest,
				nil,
				"El precio es requerido",
			))
			return
		}

		productTemporal, err := prod.service.UpdateByID(int(id), request.Name, request.Color, request.Code, request.Price)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(
				http.StatusBadRequest,
				nil,
				err.Error(),
			))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(
			http.StatusOK,
			productTemporal,
			"",
		))
	}
}
