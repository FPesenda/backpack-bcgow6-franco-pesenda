package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/warehouse"
	"github.com/stretchr/testify/assert"
)

type responseWarehouseTest struct {
	Data domain.Warehouse `json:"data"`
}
type responseWarehousesTest struct {
	Data []domain.Warehouse `json:"data"`
}
type errorWarehouseResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

var testData []domain.Warehouse = []domain.Warehouse{
	{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	},
	{
		ID:                 2,
		Address:            "sarasa2",
		Telephone:          "114292132131",
		WarehouseCode:      "test2",
		MinimumCapacity:    10,
		MinimumTemperature: 10},
}
var testDataDelete []domain.Warehouse = []domain.Warehouse{
	{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	},
	{
		ID:                 2,
		Address:            "sarasa2",
		Telephone:          "114292132131",
		WarehouseCode:      "test2",
		MinimumCapacity:    10,
		MinimumTemperature: 10},
}
var testDataEmpty []domain.Warehouse = []domain.Warehouse{}

func TestWarehouseHandlerGetallHappy(t *testing.T) {
	//arrange
	objtResp := responseWarehousesTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodGet, "/warehouses", nil)
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objtResp)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, testData, objtResp.Data)
}
func TestWarehouseHandlerGetallSad(t *testing.T) {
	//arrange
	expectedErr := "error: repository down"
	repoError := errors.New("repository down")
	objtResp := errorResponse{}
	serviceMock := warehouse.ServiceMock{ErrRepo: repoError}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodGet, "/warehouses", nil)
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objtResp)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetAllFlag)
	assert.Equal(t, expectedErr, objtResp.Message)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
func TestWarehouseHandlerGetallDBempty(t *testing.T) {
	//arrange
	expInfoMsg := "the data base has no records"
	objRes := map[string]string{}
	serviceMock := warehouse.ServiceMock{DBDummy: testDataEmpty}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodGet, "/warehouses", nil)
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetAllFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expInfoMsg, objRes["data"])
}

func TestWarehouseHandlerGetByIdHappy(t *testing.T) {
	//arrange
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodGet, "/warehouses/1", nil)
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, testData[0], objetResp.Data)
}
func TestWarehouseHandlerGetByIdSad(t *testing.T) {
	//arrange
	expectedErr := "error: notFound"
	repoError := errors.New("notFound")
	objtResp := errorResponse{}
	serviceMock := warehouse.ServiceMock{ErrNotFound: repoError}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodGet, "/warehouses/2", nil)
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objtResp)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetFlag)
	assert.Equal(t, expectedErr, objtResp.Message)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
func TestWarehouseHandlerCreateHappy(t *testing.T) {
	//arrange
	expectedDB := testData
	expectedDB = append(expectedDB, domain.Warehouse{
		ID:                 3,
		Address:            "calle Pesto",
		Telephone:          "1233443",
		WarehouseCode:      "AAASQ",
		MinimumCapacity:    10,
		MinimumTemperature: 15,
	})
	body := `{"address": "calle Pesto","telephone": "1233443","warehouse_code": "AAASQ","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPost, "/warehouses", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.True(t, serviceMock.CreateFlag)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
}

func TestWarehouseHandlerCreateWithBadAdress(t *testing.T) {
	//arrange
	expectedDB := testData
	body := `{"address": "","telephone": "1233443","warehouse_code": "AAASQ","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPost, "/warehouses", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.False(t, serviceMock.CreateFlag)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
}
func TestWarehouseHandlerCreateWithBadTelephone(t *testing.T) {
	//arrange
	expectedDB := testData
	body := `{"address": "sarasa","telephone": "","warehouse_code": "AAASQ","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPost, "/warehouses", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.False(t, serviceMock.CreateFlag)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
}

func TestWarehouseHandlerCreateWithBadWarehouseCode(t *testing.T) {
	//arrange
	expectedDB := testData
	body := `{"address": "sarasa","telephone": "21321","warehouse_code": "","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPost, "/warehouses", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.False(t, serviceMock.CreateFlag)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
}

func TestWarehouseHandlerCreateWithSameWHCode(t *testing.T) {
	//arrange
	expectedDB := testData
	body := `{"address": "","telephone": "1233443","warehouse_code": "test1","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPost, "/warehouses", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.False(t, serviceMock.CreateFlag)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
}

func TestWarehouseHandlerUpdateHappy(t *testing.T) {
	//arrange
	expectedDB := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "sarasa3",
			Telephone:          "1233443",
			WarehouseCode:      "AAASQ",
			MinimumCapacity:    10,
			MinimumTemperature: 15,
		},
		{
			ID:                 2,
			Address:            "sarasa2",
			Telephone:          "114292132131",
			WarehouseCode:      "test2",
			MinimumCapacity:    10,
			MinimumTemperature: 10},
	}
	expectedUpdated := domain.Warehouse{
		ID:                 1,
		Address:            "sarasa3",
		Telephone:          "1233443",
		WarehouseCode:      "AAASQ",
		MinimumCapacity:    10,
		MinimumTemperature: 15,
	}
	body := `{"address":"sarasa3","telephone": "1233443","warehouse_code": "AAASQ","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData, DataDummy: testData[0]}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPatch, "/warehouses/1", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.UpdateFlag)
	assert.Equal(t, expectedUpdated, objetResp.Data)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestWarehouseHandlerUpdateNotFoundSad(t *testing.T) {
	//assert
	expectedDB := testData
	body := `{"telephone": "1233443","warehouse_code": "AAASQ","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData, DataDummy: testData[0]}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPatch, "/warehouses/5", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.UpdateFlag)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
	assert.Equal(t, http.StatusNotFound, rr.Code)

}
func TestWarehouseHandlerUpdateSameWHCodeSad(t *testing.T) {
	//assert
	expectedDB := testData
	body := `{"telephone": "1233443","warehouse_code": "test2","minimum_capacity": 10,"minimum_temperature": 15}`
	objetResp := responseWarehouseTest{}
	serviceMock := warehouse.ServiceMock{DBDummy: testData, DataDummy: testData[0]}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodPatch, "/warehouses/1", []byte(body))
	router.ServeHTTP(rr, req)
	//act
	err := json.Unmarshal(rr.Body.Bytes(), &objetResp)
	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.UpdateFlag)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
	assert.Equal(t, http.StatusConflict, rr.Code)

}

func TestWarehouseHandlerDeleteHappy(t *testing.T) {
	//arrange
	expectedDB := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "sarasa",
			Telephone:          "1142910831",
			WarehouseCode:      "test1",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		}}
	serviceMock := warehouse.ServiceMock{DBDummy: testDataDelete}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodDelete, "/warehouses/2", nil)
	//act
	router.ServeHTTP(rr, req)
	//assert
	assert.True(t, serviceMock.DeleteFlag)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
	assert.Equal(t, http.StatusNoContent, rr.Code)
}
func TestWarehouseHandlerDeleteSad(t *testing.T) {
	//arrange
	expectedDB := testDataDelete
	serviceMock := warehouse.ServiceMock{DBDummy: testDataDelete}
	handler := NewWarehouse(&serviceMock)
	router := CreateServer(handler, "/warehouses")
	req, rr := CreateRequestTest(http.MethodDelete, "/warehouses/5", nil)
	//act
	router.ServeHTTP(rr, req)
	//assert
	assert.False(t, serviceMock.DeleteFlag)
	assert.Equal(t, expectedDB, serviceMock.DBDummy)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
