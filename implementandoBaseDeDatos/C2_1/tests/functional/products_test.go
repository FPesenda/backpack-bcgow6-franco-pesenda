package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/cmd/server/handler"
	cnn "github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/db"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/domains"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var s = createServer()

type Response struct {
	Data  domains.Product `json:"data,omitempty"`
	Error string          `json:"error,omitempty"`
	Code  int             `json:"code"`
}

type ResponseSlice struct {
	Data  []domains.Product `json:"data,omitempty"`
	Error string            `json:"error,omitempty"`
	Code  int               `json:"code"`
}

type ErrResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"code"`
}

func createServer() *gin.Engine {
	os.Setenv("USERNAME", "root")
	os.Setenv("PASSWORD", "")
	os.Setenv("DATABASE", "storage")

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)

	p := handler.NewProduct(serv)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetByName())
	pr.POST("/", p.Store())
	pr.GET("", p.GetAll())
	pr.DELETE("/:id", handler.IdValidationMiddleWare(), p.Delete())

	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestStoreProduct_Ok(t *testing.T) {
	new := domains.Product{
		Name:  "producto nuevo",
		Type:  "producto tipo",
		Count: 3,
		Price: 84.4,
	}

	product, err := json.Marshal(new)
	require.Nil(t, err)

	req, rr := createRequest(http.MethodPost, "/api/v1/products/", string(product))
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusCreated, rr.Code)
	// struct for assertion
	var p Response
	err = json.Unmarshal(rr.Body.Bytes(), &p)
	require.Nil(t, err)

	new.ID = p.Data.ID
	assert.Equal(t, new, p.Data)
}

func TestGetByNameProduct_Ok(t *testing.T) {
	req, rr := createRequest(http.MethodGet, "/api/v1/products/", `{"nombre":"producto nuevo"}`)
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAllHappy(t *testing.T) {
	request, response := createRequest(http.MethodGet, "/api/v1/products", "")
	s.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDelete(t *testing.T) {
	//CREAR PRODUCTO
	new := domains.Product{
		Name:  "nuevito",
		Type:  "producto tipo",
		Count: 3,
		Price: 84.4,
	}

	product, err := json.Marshal(new)
	require.Nil(t, err)

	req, resp := createRequest(http.MethodPost, "/api/v1/products/", string(product))
	s.ServeHTTP(resp, req)

	var p Response
	err = json.Unmarshal(resp.Body.Bytes(), &p)

	t.Run("delete ok", func(t *testing.T) {
		rr := ""
		request, response := createRequest(http.MethodDelete, fmt.Sprint("/api/v1/products/", p.Data.ID), "")
		s.ServeHTTP(response, request)
		json.Unmarshal(response.Body.Bytes(), &rr)
		assert.Equal(t, http.StatusNoContent, response.Code)
	})
	t.Run("serch deleted prouct ok ", func(t *testing.T) {
		req, rr := createRequest(http.MethodGet, "/api/v1/products/", fmt.Sprint(`{"nombre":"`, p.Data.Name, `"}`))
		s.ServeHTTP(rr, req)

		// assert code
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
	t.Run("serch deleted product in all registers", func(t *testing.T) {
		request, response := createRequest(http.MethodGet, "/api/v1/products", "")
		s.ServeHTTP(response, request)
		var ps ResponseSlice

		err := json.Unmarshal(response.Body.Bytes(), &ps)
		assert.Nil(t, err)
		deleted := true
		for _, v := range ps.Data {
			if p.Data.ID == v.ID {
				deleted = false
			}
		}
		assert.True(t, deleted)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
