package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
	"github.com/gin-gonic/gin"
)

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

// ListBuyer godoc
// @Summary     Buyer by id
// @Tags        Buyer
// @Description get Buyer by id
// @Produce     json
// @Param       token header   string true "token"
// @Param       id    path     int    true "Buyer id"
// @Success     200   {object} web.response
// @Failure     400   {object} web.response
// @Failure     401   {object} web.response
// @Failure     404   {object} web.response
// @Router      /buyers/{id} [get]
func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.GetInt("id_validated")
		p, err := b.buyerService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Response(c, http.StatusOK, p)

	}
}

// ListBuyers godoc
// @Summary     List Buyers
// @Tags        Buyer
// @Description get Buyers
// @Produce     json
// @Param       token header   string true "token"
// @Success     200   {object} web.response
// @Failure     404   {object} web.response
// @Router      /buyers [get]
func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := b.buyerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusNotFound, "the data base has no records")
			return
		}

		web.Success(c, http.StatusOK, p)
	}
}

// StoreProducts godoc
// @Summary     Store Buyers
// @Tags        Buyer
// @Description store buyers
// @Accept      json
// @Param       token header   string            true "token"
// @Param       buyer body     request.BuyerPost true "Buyer to store"
// @Success     201   {object} web.response
// @Failure     400   {object} web.response
// @Failure     401   {object} web.response
// @Failure     404   {object} web.response
// @Failure     409   {object} web.response
// @Failure     422   {object} web.response
// @Router      /buyers [post]
func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.BuyerPost
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error : %s", err)
			return
		}

		buyer := domain.Buyer{ID: 0, CardNumberID: req.CardNumberID, FirstName: req.FirstName, LastName: req.LastName}

		newId, err := b.buyerService.Save(c, buyer)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		buyer.ID = newId

		web.Success(c, http.StatusCreated, buyer)
	}
}

// UpdateBuyer godocc
// @Summary     Update Buyer
// @Tags        Buyer
// @Description update Buyer
// @Accept      json
// @Param       token   header   string             true "token"
// @Param       id      path     int                true "Buyer id"
// @Param       request body     request.BuyerPatch true "Buyer to update"
// @Success     200     {object} web.response
// @Failure     400     {object} web.response
// @Failure     401     {object} web.response
// @Failure     404     {object} web.response
// @Failure     422     {object} web.response
// @Router      /buyers/{id} [PATCH]
func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.BuyerPatch
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error %v", err)
			return
		}

		id := c.GetInt("id_validated")

		buyer, err := b.buyerService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error : buyer not found %d", id)
			return
		}

		if req.FirstName != "" {
			buyer.FirstName = req.FirstName
		}
		if req.LastName != "" {
			buyer.LastName = req.LastName
		}

		err = b.buyerService.Update(c, buyer)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, buyer)
	}

}

// DeleteBuyer godoc
// @Summary     Delete Buyer
// @Tags        Buyer
// @Description delete Buyer
// @Produce     json
// @Param       token header   string true "token"
// @Param       id    path     int    true "Buyer id"
// @Success     204   {object} web.response
// @Failure     400   {object} web.response
// @Failure     401   {object} web.response
// @Failure     404   {object} web.response
// @Router      /buyers/{id} [delete]
func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		err := b.buyerService.Delete(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusNoContent, nil)

	}

}
