package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/domains"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/products"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/pkg/request"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestName struct {
	Name string `json:"nombre" binding:"required"`
}

type Product struct {
	service products.Service
}

var (
	NOT_FOUND = "producto no encontrado"
)

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (s *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.Product
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product := domains.Product(domains.Product{
			Name:  req.Name,
			Type:  req.Type,
			Count: req.Count,
			Price: req.Price,
		})
		id, err := s.service.Store(product)
		if err != nil {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		product.ID = id
		c.JSON(http.StatusCreated, web.NewResponse(product, "", http.StatusCreated))
	}
}

func (s *Product) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestName
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product, err := s.service.GetByName(req.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(product, "", http.StatusOK))
	}
}

func (s *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := s.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(products, "", http.StatusOK))
	}
}

func (s *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product request.Product

		id := ctx.GetInt("id_validated")

		if !s.service.Exist(id) {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, fmt.Sprint("Producto con id ", id, " no enontrado"), http.StatusNotFound))
			return
		}

		err := ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}
		productNew := domains.Product{
			ID:    id,
			Name:  product.Name,
			Type:  product.Type,
			Count: product.Count,
			Price: product.Price,
		}
		err = s.service.Update(productNew)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(productNew, "", http.StatusOK))
	}
}

func (s *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id_validated")
		if !s.service.Exist(id) {
			ctx.JSON(http.StatusNotFound, web.NewResponse(nil, fmt.Sprint(NOT_FOUND, " ", id), http.StatusNotFound))
			return
		}

		err := s.service.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(nil, "", http.StatusNoContent))

	}
}
