package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-testing/test/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockService mock.ServiceMock) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handler := NewHandler(&mockService)

	r := gin.Default()

	pr := r.Group("/api/v1/products")
	{
		pr.GET("", handler.GetProducts)
	}

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetProductsHappy(t *testing.T) {
	//arange
	data := []domain.Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	var response []domain.Product
	mockService := mock.ServiceMock{
		Data: data,
	}
	router := createServer(mockService)
	request, rr := createRequestTest(http.MethodGet, "/api/v1/products?seller_id="+data[0].SellerID, "")
	//act
	router.ServeHTTP(rr, request)
	//assert
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, mockService.Data, response)
}

func TestGetProductsSadEmptyId(t *testing.T) {
	//arange
	mockService := mock.ServiceMock{}
	var response map[string]string
	expectedError := map[string]string{
		"error": "seller_id query param is required",
	}
	router := createServer(mockService)
	request, rr := createRequestTest(http.MethodGet, `/api/v1/products?seller_id=`, "")
	//act
	router.ServeHTTP(rr, request)
	//assert
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expectedError, response)
}

func TestGetProductsSadServerFail(t *testing.T) {
	//arange
	mockService := mock.ServiceMock{
		ErrRepository: errors.New("Error"),
	}
	var response map[string]string
	expectedError := map[string]string{
		"error": "Error",
	}
	router := createServer(mockService)
	request, rr := createRequestTest(http.MethodGet, `/api/v1/products?seller_id=FEX112AC`, "")
	//act
	router.ServeHTTP(rr, request)
	//assert
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, expectedError, response)
}
