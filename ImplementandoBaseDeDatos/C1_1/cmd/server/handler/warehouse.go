package handler

import (
	"fmt"
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

// ListWarehouse godoc
// @Summary     Warehouse by id
// @Tags        Warehouse
// @Description get Warehouse by id
// @Produce     json
// @Param       token header   string true "token"
// @Param       id    path     int    true "Product id"
// @Success     200   {object} web.response
// @failure     400   {object} web.response
// @failure     404   {object} web.response
// @Router      /warehouses/{id} [get]
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		p, err := w.warehouseService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err)
			return
		}
		web.Success(c, http.StatusOK, p)
	}

}

// ListWarehouses godoc
// @Summary     List Warehouses
// @Tags        Warehouse
// @Description get Warehouses
// @Produce     json
// @Param       token header   string true "token"
// @Success     200   {object} web.response
// @failure     404   {object} web.response
// @Router      /warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := w.warehouseService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err)
			return
		}
		if data == nil {
			web.Success(c, http.StatusOK, "the data base has no records")
			return
		}
		web.Success(c, http.StatusOK, data)
	}
}

// StoreWarehouses godoc
// @Summary     Store Warehouses
// @Tags        Warehouse
// @Description store warehouses
// @Accept      json
// @Param       token     header   string            true "token"
// @Param       warehouse body     request.Warehouse true "Warehouse to store"
// @Success     201       {object} web.response
// @failure     404       {object} web.response
// @failure     422       {object} web.response
// @failure     409       {object} web.response
// @Router      /warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wh request.Warehouse
		if err := c.Bind(&wh); err != nil {
			web.Error(c, http.StatusNotFound, "error: %s ", err)
			return
		}
		if w.warehouseService.Exists(c, wh.WarehouseCode) {
			web.Error(c, http.StatusConflict, "error : warehouse code already in use")
			return
		}
		if wh.WarehouseCode == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error : bad warehouseCode")
			return
		}
		if wh.Address == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error : bad Adress")
			return
		}
		if wh.Telephone == "" {
			web.Error(c, http.StatusUnprocessableEntity, "error : bad Telephone")
			return
		}
		if wh.MinimumCapacity < 0 {
			web.Error(c, http.StatusUnprocessableEntity, "error : capacity can't be less than 0")
		}
		if wh.MinimumCapacity == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "error : bad Capacity")
			return
		}
		if wh.MinimumTemperature == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "error : bad Temperature")
			return
		}

		createdID, err := w.warehouseService.Save(c, domain.Warehouse{
			Address:            wh.Address,
			Telephone:          wh.Telephone,
			WarehouseCode:      wh.WarehouseCode,
			MinimumCapacity:    wh.MinimumCapacity,
			MinimumTemperature: wh.MinimumTemperature,
		})

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error: %s", err)
			return
		}
		created, err := w.warehouseService.Get(c, createdID)
		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error: %s", err)
			return

		}
		web.Success(c, http.StatusCreated, created)
	}
}

// UpdateWarehouse godocc
// @Summary     Update Warehouse
// @Tags        Warehouse
// @Description update Warehouse
// @Accept      json
// @Param       token   header   string            true "token"
// @Param       id      path     int               true "Warehouse id"
// @Param       request body     request.Warehouse true "Warehouse to update"
// @Success     200     {object} web.response
// @failure     400     {object} web.response
// @failure     404     {object} web.response
// @failure     409     {object} web.response
// @Router      /warehouses/{id} [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var whReq request.Warehouse
		id := c.GetInt("id_validated")
		if err := c.ShouldBindJSON(&whReq); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err)
			return

		}
		whBD, errGet := w.warehouseService.Get(c, int(id))
		if errGet != nil {
			web.Error(c, http.StatusNotFound, "error: %s", errGet)
			return
		}
		if whReq.WarehouseCode != "" {
			if w.warehouseService.Exists(c, whReq.WarehouseCode) && whReq.WarehouseCode != whBD.WarehouseCode {
				web.Error(c, http.StatusConflict, "error : warehouse code already in use")
				return
			} else {
				whBD.WarehouseCode = whReq.WarehouseCode
			}
		}
		if whReq.Address != "" {
			whBD.Address = whReq.Address
		}
		if whReq.Telephone != "" {
			whBD.Telephone = whReq.Telephone
		}
		if whReq.MinimumCapacity < 0 {
			web.Error(c, http.StatusUnprocessableEntity, "error : capacity can't be less than 0")
		}
		if whReq.MinimumCapacity != 0 {
			whBD.MinimumCapacity = whReq.MinimumCapacity
		}
		if whReq.MinimumTemperature != 0 {
			whBD.MinimumTemperature = whReq.MinimumTemperature
		}
		err := w.warehouseService.Update(c, whBD)
		if err != nil {
			web.Error(c, http.StatusNotFound, fmt.Sprint(err))
		}
		web.Success(c, http.StatusOK, whBD)
	}
}

// DeleteWarehouse godoc
// @Summary     Delete Warehouse
// @Tags        Warehouse
// @Description delete Warehouse
// @Produce     json
// @Param       token header   string true "token"
// @Param       id    path     int    true "Warehouse id"
// @Success     204   {object} web.response
// @failure     400   {object} web.response
// @failure     404   {object} web.response
// @Router      /warehouses/{id} [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		whBD, errGet := w.warehouseService.Get(c, int(id))
		if errGet != nil {
			web.Error(c, http.StatusNotFound, "error: %s", errGet)
			return
		}
		err := w.warehouseService.Delete(c, whBD.ID)
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s ", err)
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}
