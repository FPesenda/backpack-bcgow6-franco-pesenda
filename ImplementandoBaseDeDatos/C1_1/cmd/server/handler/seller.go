package handler

import (
	"fmt"
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
	"github.com/gin-gonic/gin"
)

type Request struct {
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

// ListSeller godoc
// @Summary List Sellers
// @Tags Sellers
// @Description get Sellers
// @Produce  json
// @Success 200 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Router /sellers [get]
// @Param token header string true "token"
func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error(), "")
			return
		}
		if sellers == nil {
			web.Success(c, 200, "the data base has no records")
			return
		}
		web.Success(c, 200, sellers)

	}
}

// GetAllSeller godoc
// @Summary Seller by id
// @Tags Sellers
// @Description get Seller by id
// @Produce  json
// @Param id path int true "Seller id"
// @Success 200 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Router /sellers/{id} [get]
// @Param token header string true "token"
func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		seller, err := s.sellerService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "seller with %d does not exist ", id)
			return
		}
		web.Success(c, http.StatusOK, seller)
	}
}

// CreateSeller godoc
// @Summary Create Seller
// @Tags Sellers
// @Description create Seller
// @Accept  json
// @Produce  json
// @Param seller body request.RequestSeller true "Seller to create"
// @Success 201 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Router /sellers [post]
// @Param token header string true "token"
func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "error : not a json")
			return
		}
		if req.Address == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error : the address is required")
			return
		}
		if req.CompanyName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error : the company's name is required")
			return
		}
		if req.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error : the telephone is required")
			return
		}
		if req.CID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "error : the CID is required")
			return
		}
		seller, err := s.sellerService.Create(c, domain.Seller{
			CID:         req.CID,
			CompanyName: req.CompanyName,
			Address:     req.Address,
			Telephone:   req.Telephone,
		})
		if err != nil {
			web.Error(c, http.StatusConflict, "seller with CID %d already exists", req.CID)
			return
		}
		web.Success(c, http.StatusCreated, seller)

	}
}

// UpdateSeller godocc
// @Summary Update Seller
// @Tags Sellers
// @Description update Seller
// @Accept  json
// @Produce  json
// @Param id path int true "Seller id"
// @Param request body request.RequestSeller true "Seller to update"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 422 {object} web.errorResponse
// @Router /sellers/{id} [PATCH]
// @Param token header string true "token"
func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.RequestSeller
		id := c.GetInt("id_validated")
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error: %s", err)
			return
		}
		fmt.Printf("%+v", req)
		seller, err := s.sellerService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "seller with %d does not exist ", id)
			return
		}

		if req.CID != 0 {
			if s.sellerService.Exist(c, req.CID) && req.CID != seller.CID {
				web.Error(c, http.StatusConflict, "seller with CID %d already exists ", req.CID)
				return
			}
			seller.CID = req.CID
		}
		if req.Address != "" {
			seller.Address = req.Address
		}
		if req.CompanyName != "" {
			seller.CompanyName = req.CompanyName
		}
		if req.Telephone != "" {
			seller.Telephone = req.Telephone
		}
		updated, err := s.sellerService.Update(c, seller)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusOK, updated)
	}
}

// DeleteSeller godoc
// @Summary Delete Seller
// @Tags Sellers
// @Description delete Seller
// @Produce  json
// @Param id path int true "Seller id"
// @Success 204 {object} web.response
// @Failure 404 {object} web.errorResponse
// @Router /sellers/{id} [delete]
// @Param token header string true "token"
func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		err := s.sellerService.Delete(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "seller with %d does not exist ", id)
			return
		}

		web.Response(c, http.StatusNoContent, "")
	}
}
