package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/cmd/server/handler"
	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/domain"
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
		productsRouter.POST("/", product.Store())
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

func updateHappyTest(t *testing.T) {
	//ARANGE
	mockStorage := mocks.MockStorage{
		Data: []domain.Product{
			{
				ID:    1,
				Name:  "Product1",
				Type:  "Test",
				Count: 10,
				Price: 10.5,
			},
		},
		ErrWrite: "",
		ErrRead:  "",
	}
	expectedProduct := domain.Product{
		ID:    1,
		Name:  "PRD 1 Updated",
		Type:  "Type Updated",
		Count: 10,
		Price: 10.5,
	}
	router := createServer(mockStorage)
	request, rr := createRequestTest(http.MethodPost, "/products/", `{
        "Name": "PRD 1 Updated",
		"Type": "Type Updated",
		"Count": 99,
		"Price": 9.9
    }`)
	//ACT
	router.ServeHTTP(rr, request)
	//ASSERT
	err := json.Unmarshal(rr.Body.Bytes(), &request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, expectedProduct, rr.Body)
}
