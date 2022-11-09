package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/product"
	mocks "github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/products"
	"github.com/stretchr/testify/assert"
)

type responseStringTest struct {
	Data string `json:"data"`
}
type responseProductTest struct {
	Data domain.Product `json:"data"`
}

type responseProductsTest struct {
	Data []domain.Product `json:"data"`
}

var path string = "/products"

func TestGetAllHappy(t *testing.T) {
	//Arange
	expectedData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "Descripciñon 2",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	service := mocks.ServiceProductMock{
		Data: expectedData,
	}
	response := responseProductsTest{}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, resp := CreateRequestTest(http.MethodGet, path, nil)
	//Act
	router.ServeHTTP(resp, request)
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expectedData, response.Data)
}

func TestGetAllSad(t *testing.T) {
	//Arange
	expectedError := "error Get All Error"
	service := mocks.ServiceProductMock{
		Err: errors.New("Get All Error"),
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodGet, path, nil)
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestGetAllSadEmpty(t *testing.T) {
	//Arange
	expectedData := "the data base has no records"
	service := mocks.ServiceProductMock{
		Data: nil,
	}
	response := responseStringTest{}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, resp := CreateRequestTest(http.MethodGet, path, nil)
	//Act
	router.ServeHTTP(resp, request)
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	fmt.Println("DATA COONSEGUIDA: ", resp.Body)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expectedData, response.Data)
}

func TestGetById(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := domain.Product{
		ID:             1,
		Description:    "Descripciñon 1",
		ExpirationRate: 10,
		FreezingRate:   20,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	idToSerch := 1
	service := mocks.ServiceProductMock{
		Data: initialData,
	}
	response := responseProductTest{}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, resp := CreateRequestTest(http.MethodGet, fmt.Sprint(path, "/", idToSerch), nil)
	//Act
	router.ServeHTTP(resp, request)
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expectedData, response.Data)
}

func TestGetByIdSad(t *testing.T) {
	//Arange
	expectedError := "error Get By Id Error"
	service := mocks.ServiceProductMock{
		Err:               errors.New("Get By Id Error"),
		GetByIdIdExecuted: false,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodGet, fmt.Sprint(path, "/1"), nil)
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.True(t, service.GetByIdIdExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestGetByIdSadNonExistent(t *testing.T) {
	//Arange
	expectedError := "error product not found"
	service := mocks.ServiceProductMock{
		ErrFindById:       product.ErrNotFound,
		GetByIdIdExecuted: false,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodGet, fmt.Sprint(path, "/1"), nil)
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.True(t, service.GetByIdIdExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestCreateHappy(t *testing.T) {
	//Arange
	initialData := []domain.Product{}
	expectedData := domain.Product{
		ID:             1,
		Description:    "Mayonesa",
		ExpirationRate: 10,
		FreezingRate:   20,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	body := `{
		"description":"Mayonesa",
		"expiration_rate":10,
		"freezing_rate":20,
		"height": 10.1,
		"length":20.1,
		"netweight":30.1,
		"product_code":"CODE1",
		"recommended_freezing_temperature":40.1,
		"width": 50.1,
		"product_type_id":1,
		"seller_id":1
	}`
	response := responseProductTest{}
	service := mocks.ServiceProductMock{
		Data: initialData,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPost, path, []byte(body))
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, respons.Code)
	assert.Equal(t, expectedData, response.Data)
}

func TestCreateSad(t *testing.T) {
	//Arange
	expectedError := "error Save Error"
	service := mocks.ServiceProductMock{
		Err:            errors.New("Save Error"),
		SaveIsExecuted: false,
	}
	body := `{
		"description":"Mayonesa",
		"expiration_rate":10,
		"freezing_rate":20,
		"height": 10.1,
		"length":20.1,
		"netweight":30.1,
		"product_code":"CODE1",
		"recommended_freezing_temperature":40.1,
		"width": 50.1,
		"product_type_id":1,
		"seller_id":1
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPost, path, []byte(body))
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.True(t, service.SaveIsExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestCreateSadUnprocesableEntity(t *testing.T) {
	//Arange
	expectedError := "error: json: cannot unmarshal string into Go struct field Product.height of type float32"
	service := mocks.ServiceProductMock{
		Err:            errors.New("error: json: cannot unmarshal string into Go struct field Product.height of type float32"),
		SaveIsExecuted: false,
	}
	body := `{
		"description":"Mayonesa",
		"expiration_rate":10,
		"freezing_rate":20,
		"height": "ERROR ACA",
		"length":20.1,
		"netweight":30.1,
		"product_code":"CODE1",
		"recommended_freezing_temperature":40.1,
		"width": 50.1,
		"product_type_id":1,
		"seller_id":1
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPost, path, []byte(body))
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.False(t, service.SaveIsExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestCreateSadSaveExistinProductInBDD(t *testing.T) {
	//Arange
	expectedData := "error fail in create new product"
	response := errorResponse{}
	service := mocks.ServiceProductMock{
		ErrSave: product.ErrSaveFalied,
	}
	body := `{
		"description":"Mayonesa",
		"expiration_rate":10,
		"freezing_rate":20,
		"height": 10.1,
		"length":20.1,
		"netweight":30.1,
		"product_code":"CODE1",
		"recommended_freezing_temperature":40.1,
		"width": 50.1,
		"product_type_id":1,
		"seller_id":1
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPost, path, []byte(body))
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, respons.Code)
	assert.Equal(t, expectedData, response.Message)
}

func TestCreateSadNeedAllFields(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Mayonesa",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := "error: Key: 'Product.ExpirationRate' Error:Field validation for 'ExpirationRate' failed on the 'required' tag"
	body := `{
		"description":"Mayonesa",
		"freezing_rate":20,
		"height": 10.1,
		"length":20.1,
		"netweight":30.1,
		"product_code":"CODE1",
		"recommended_freezing_temperature":40.1,
		"width": 50.1,
		"product_type_id":1,
		"seller_id":1
	}`
	response := errorResponse{}
	service := mocks.ServiceProductMock{
		Data: initialData,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPost, path, []byte(body))
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, respons.Code)
	assert.Equal(t, expectedData, response.Message)
}

func TestUpdateHappyFull(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Mayonesa",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := domain.Product{
		ID:             1,
		Description:    "Mayonesa",
		ExpirationRate: 10,
		FreezingRate:   99,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	idToUpdate := 1
	body := `{
		"freezing_rate":99
	}`
	response := responseProductTest{}
	service := mocks.ServiceProductMock{
		Data:             initialData,
		UpdateIsExecuted: false,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPatch, fmt.Sprint(path, "/", idToUpdate), []byte(body))
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.True(t, service.UpdateIsExecuted)
	assert.Equal(t, http.StatusOK, respons.Code)
	assert.Equal(t, expectedData, response.Data)
}

func TestUpdateHappy(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Mayonesa",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := domain.Product{
		ID:             1,
		Description:    "Mayoneson",
		ExpirationRate: 10,
		FreezingRate:   99,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	idToUpdate := 1
	body := `{
		"description": "Mayoneson",
		"expiration_rate": 10,
		"freezing_rate": 99,
		"height": 10.1,
		"length": 20.1,
		"netweight": 30.1,
		"product_code": "CODE1",
		"recommended_freezing_temperature": 40.1,
		"width": 50.1,
		"product_type_id": 1,
		"seller_id": 1
	}`
	response := responseProductTest{}
	service := mocks.ServiceProductMock{
		Data:             initialData,
		UpdateIsExecuted: false,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPatch, fmt.Sprint(path, "/", idToUpdate), []byte(body))
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.True(t, service.UpdateIsExecuted)
	assert.Equal(t, http.StatusOK, respons.Code)
	assert.Equal(t, expectedData, response.Data)
}

func TestUpdateSad(t *testing.T) {
	//Arange
	expectedError := "error update"
	service := mocks.ServiceProductMock{
		ErrUpdate:        errors.New("update"),
		UpdateIsExecuted: false,
	}
	body := `{
		"freezing_rate":99
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPatch, fmt.Sprint(path, "/1"), []byte(body))
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.True(t, service.UpdateIsExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestUpdateSadNotFound(t *testing.T) {
	//Arange
	expectedError := "error product not found"
	service := mocks.ServiceProductMock{
		Err:              errors.New("not found"),
		UpdateIsExecuted: false,
	}
	body := `{
		"freezing_rate":99
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPatch, fmt.Sprint(path, "/1"), []byte(body))
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.False(t, service.UpdateIsExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestUpdateSadUnprocesableEntity(t *testing.T) {
	//Arange
	expectedError := "error: json: cannot unmarshal string into Go struct field ProductUpdate.freezing_rate of type int"
	service := mocks.ServiceProductMock{
		Err:              errors.New("error: json: cannot unmarshal string into Go struct field ProductUpdate.freezing_rate of type int"),
		UpdateIsExecuted: false,
	}
	body := `{
		"freezing_rate":"ERROR"
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPatch, fmt.Sprint(path, "/1"), []byte(body))
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.False(t, service.UpdateIsExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestUpdateSadProductCodeExist(t *testing.T) {
	//Arange
	expectedError := "error the resource alredy exist in the database"
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Mayonesa",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "Mayonesa",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE2",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	service := mocks.ServiceProductMock{
		Data:             initialData,
		UpdateIsExecuted: false,
	}
	body := `{
		"freezing_rate":99,
		"product_code": "CODE1"
	}`
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodPatch, fmt.Sprint(path, "/2"), []byte(body))
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.False(t, service.UpdateIsExecuted)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestDeleteHappy(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Mayonesa",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	service := mocks.ServiceProductMock{
		Data: initialData,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodDelete, fmt.Sprint(path, "/1"), nil)
	//Act
	router.ServeHTTP(respons, request)
	//Assert
	assert.Equal(t, http.StatusNoContent, respons.Code)
}

func TestDeleteSad(t *testing.T) {
	//Arange
	expectedError := "error delete"
	service := mocks.ServiceProductMock{
		Err: errors.New("delete"),
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodDelete, fmt.Sprint(path, "/1"), nil)
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}

func TestDeleteSadIdNotFound(t *testing.T) {
	//Arange
	expectedError := "error product not found"
	service := mocks.ServiceProductMock{
		ErrDelete: product.ErrNotFound,
	}
	handler := NewProduct(&service)
	router := CreateServer(handler, path)
	request, respons := CreateRequestTest(http.MethodDelete, fmt.Sprint(path, "/1"), nil)
	response := errorResponse{}
	//Act
	router.ServeHTTP(respons, request)
	err := json.Unmarshal(respons.Body.Bytes(), &response)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, respons.Code)
	assert.Equal(t, expectedError, response.Message)
}
