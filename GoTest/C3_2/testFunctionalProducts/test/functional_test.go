package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/cmd/server/handler"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/products"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mocks.MockStorage) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repository := products.NewRepository(&mockStore)
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	r := gin.Default()

	productsRouter := r.Group("/products")
	{
		productsRouter.GET("/", product.GetAll())
		productsRouter.POST("/:id", product.Store())
		productsRouter.DELETE("/:id", product.Delete())
	}

	return r

}

func createRequestTest(method, url, body string) (request *http.Request, response *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	request = req
	response = httptest.NewRecorder()
	return
}

func TestHappyUpdate(t *testing.T) {
	//ARANGE
	mockStorage := mocks.MockStorage{
		Data: []products.Products{
			{
				Id:    1,
				Name:  "Product1",
				Color: "Test",
				Code:  "10",
				Price: 10.5,
			},
		},
		ErrWrite: "",
		ErrRead:  "",
	}
	expectedProduct := products.Products{
		Id:    1,
		Name:  "PRD 1 Updated",
		Color: "Type Updated",
		Code:  "10",
		Price: 10.5,
	}
	router := createServer(mockStorage)
	request, rr := createRequestTest(http.MethodPost, "/products/1",
		`{
        "Name":"PRD 1 Updated",
		"Color":"Type Updated",
		"Code":"10",
		"Price":10.5
    }`)
	var response products.Products
	//ACT
	router.ServeHTTP(rr, request)
	//ASSERT
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, expectedProduct, response)
}

func TestDeleteHappy(t *testing.T) {
	//range
	data := []products.Products{
		{
			Id:    1,
			Name:  "PRD 1 Updated",
			Color: "Type Updated",
			Code:  "10",
			Price: 10.5,
		},
	}
	mockStorage := mocks.MockStorage{
		Data:     data,
		ErrWrite: "",
		ErrRead:  "",
	}
	router := createServer(mockStorage)
	request, rr := createRequestTest(http.MethodDelete, "/products/1", "")
	var response map[string]string
	expectedResponse := map[string]string{
		"data": fmt.Sprintf("El producto %d a sido eliminado correctamente", data[0].Id),
	}
	//act
	router.ServeHTTP(rr, request)
	//assert
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expectedResponse, response)
}
