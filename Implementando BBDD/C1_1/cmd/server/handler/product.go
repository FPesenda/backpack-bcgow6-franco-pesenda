package handler

import (
	"net/http"
)

type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{
		service: service,
	}
}

// ListProducts godoc
// @Summary     List Products
// @Tags        Product
// @Description get Products
// @Param       token header string true "token"
// @Produce     json
// @Success     200 {object} web.response
// @Failure     400 {object} web.errorResponse
// @Failure     404 {object} web.errorResponse
// @Router      /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productList, err := p.service.GetAll(ctx)
		if err != nil {
			managmentError(ctx, err)
			return
		}
		if productList == nil {
			web.Success(ctx, http.StatusOK, "the data base has no records")
			return
		}
		web.Success(ctx, http.StatusOK, productList)
	}
}

// GetProduct godoc
// @Summary     Product by id
// @Tags        Product
// @Description get Product by id
// @Param       token header string true "token"
// @Produce     json
// @Param       id  path     int true "Product id"
// @Success     200 {object} web.response
// @Success     400 {object} web.response
// @Failure     404 {object} web.errorResponse
// @Router      /products/{id} [get]
func (p *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id_validated")
		product, err := p.service.Get(ctx, id)
		if err != nil {
			managmentError(ctx, err)
			return
		}
		web.Success(ctx, http.StatusOK, product)
	}
}

// CreateProduct godoc
// @Summary     Create Product
// @Tags        Product
// @Description create Product
// @Param       token header string true "token"
// @Accept      json
// @Produce     json
// @Param       product body     request.Product true "Product to create"
// @Success     201     {object} web.response
// @Failure     404 {object} web.errorResponse
// @Failure     422 {object} web.errorResponse
// @Router      /products [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request request.Product
		if err := ctx.ShouldBindJSON(&request); err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "error: %s", err)
			return
		}
		productNew := domain.Product{
			Description:    request.Description,
			ExpirationRate: request.ExpirationRate,
			FreezingRate:   request.FreezingRate,
			Height:         request.Height,
			Length:         request.Length,
			Netweight:      request.Netweight,
			ProductCode:    request.ProductCode,
			RecomFreezTemp: request.RecomFreezTemp,
			Width:          request.Width,
			ProductTypeID:  request.ProductTypeID,
			SellerID:       request.SellerID,
		}
		idNewProduct, errSave := p.service.Save(ctx, productNew)
		if errSave != nil {
			managmentError(ctx, errSave)
			return
		}
		productNew.ID = idNewProduct
		web.Success(ctx, http.StatusCreated, productNew)
	}

}

// PatchProduct godocc
// @Summary     Patch Product
// @Tags        Product
// @Description patch Product
// @Param       token header string true "token"
// @Accept      json
// @Produce     json
// @Param       id      path     int                    true "Product id"
// @Param       request body     request.ProductUpdate true "Product to update"
// @Success     200     {object} web.response
// @Failure     400     {object} web.errorResponse
// @Failure     422     {object} web.errorResponse
// @Failure     404     {object} web.errorResponse
// @Router      /products/{id} [PATCH]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqt request.ProductUpdate
		id := ctx.GetInt("id_validated")
		if err := ctx.ShouldBindJSON(&reqt); err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "error: %s", err)
			return
		}
		productInBBDD, errGet := p.service.Get(ctx, id)
		if errGet != nil {
			managmentError(ctx, product.ErrNotFound)
			return
		}
		if reqt.Description != "" {
			productInBBDD.Description = reqt.Description
		}
		if reqt.ExpirationRate != 0 {
			productInBBDD.ExpirationRate = reqt.ExpirationRate
		}
		if reqt.FreezingRate != 0 {
			productInBBDD.FreezingRate = reqt.FreezingRate
		}
		if reqt.Height != 0.0 {
			productInBBDD.Height = reqt.Height
		}
		if reqt.Length != 0.0 {
			productInBBDD.Length = reqt.Length
		}
		if reqt.Netweight != 0.0 {
			productInBBDD.Netweight = reqt.Netweight
		}
		if reqt.ProductCode != "" {
			if p.service.Exist(ctx, reqt.ProductCode) && productInBBDD.ProductCode != reqt.ProductCode {
				managmentError(ctx, product.ErrProductExist)
				return
			}
			productInBBDD.ProductCode = reqt.ProductCode
		}
		if reqt.RecomFreezTemp != 0.0 {
			productInBBDD.RecomFreezTemp = reqt.RecomFreezTemp
		}
		if reqt.Width != 0.0 {
			productInBBDD.Width = reqt.Width
		}
		if reqt.ProductTypeID != 0 {
			productInBBDD.ProductTypeID = reqt.ProductTypeID
		}
		if reqt.SellerID != 0 {
			productInBBDD.SellerID = reqt.SellerID
		}
		product, errUpdate := p.service.Update(ctx, productInBBDD)
		if errUpdate != nil {
			managmentError(ctx, errUpdate)
			return
		}
		web.Success(ctx, http.StatusOK, product)
	}
}

// DeleteProduct godoc
// @Summary     Delete Product
// @Tags        Product
// @Description delete Product
// @Param       token header string true "token"
// @Produce     json
// @Param       id path int true "Product id"
// @Success     204
// @Failure     404 {object} web.errorResponse
// @Failure     400 {object} web.errorResponse
// @Router      /products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetInt("id_validated")
		err := p.service.Delete(ctx, id)
		if err != nil {
			managmentError(ctx, err)
			return
		}
		web.Success(ctx, http.StatusNoContent, "Deleted successfully")
	}
}

func managmentError(c *gin.Context, err error) {
	switch err {
	case product.ErrNotFound:
		web.Error(c, http.StatusNotFound, "error %s", err.Error())
		return
	case product.ErrSaveFalied:
		web.Error(c, http.StatusConflict, "error %s", err.Error())
		return
	case product.ErrProductExist:
		web.Error(c, http.StatusConflict, "error %s", err.Error())
		return
	default:
		web.Error(c, http.StatusBadRequest, "error %s", err.Error())
		return
	}
}
