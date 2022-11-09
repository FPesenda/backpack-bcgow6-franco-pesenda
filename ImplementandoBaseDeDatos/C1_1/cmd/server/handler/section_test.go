package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	s "github.com/extmatperez/meli_bootcamp_go_w6-4/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/section"
	"github.com/stretchr/testify/assert"
)

type responseSectionTest struct {
	Data domain.Section `json:"data"`
}
type responseSectionsTest struct {
	Data []domain.Section `json:"data"`
}

var (
	ErrRepo  error  = errors.New("repository down")
	endPoint string = "/sections"
)
var initData []domain.Section = []domain.Section{
	{
		ID:                 1,
		SectionNumber:      13,
		CurrentTemperature: 27,
		MinimumTemperature: 10,
		CurrentCapacity:    143,
		MinimumCapacity:    100,
		MaximumCapacity:    500,
		WarehouseID:        3,
		ProductTypeID:      2,
	},
	{
		ID:                 2,
		SectionNumber:      14,
		CurrentTemperature: 12,
		MinimumTemperature: 8,
		CurrentCapacity:    25,
		MinimumCapacity:    1,
		MaximumCapacity:    250,
		WarehouseID:        7,
		ProductTypeID:      16,
	},
}

// CREATE
// create a section correctly
func TestSectionHandlerCreateOk(t *testing.T) {
	//arrange
	body := `{"section_number":19,"current_temperature":24,"minimum_temperature":15,"current_capacity":372,
	"minimum_capacity":145,"maximum_capacity":550,"warehouse_id":6,"product_type_id":34}`
	objRes := responseSectionTest{}
	serviceMock := section.ServiceMock{SectionDummy: initData[0]}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPost, "/sections", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, initData[0], objRes.Data)
}

// create a section but but not all required fields are sent
func TestSectionHandlerCreateWithoutSomeFieldRequired(t *testing.T) {
	//arrange
	expErrMsg := "error Key: 'SectionPost.CurrentTemperature' Error:Field validation for 'CurrentTemperature' failed on the 'required' tag"
	body := `{"section_number":19,"minimum_temperature":15,"current_capacity":372,
	"minimum_capacity":145,"maximum_capacity":550,"warehouse_id":6,"product_type_id":34}`
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPost, "/sections", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// create a section but section number exists in the data
func TestSectionHandlerCreateAndSectionNumberAlreadyExist(t *testing.T) {
	//arrange
	expErrMsg := "error section number exists"
	body := `{"section_number":19,"current_temperature":24,"minimum_temperature":15,"current_capacity":372,
	"minimum_capacity":145,"maximum_capacity":550,"warehouse_id":6,"product_type_id":34}`
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{ErrService: s.ErrExists}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPost, "/sections", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// create a section but has problem with the repository
func TestSectionHandlerCreateAndErrRepo(t *testing.T) {
	//arrange
	expErrMsg := "error repository down"
	body := `{"section_number":19,"current_temperature":24,"minimum_temperature":15,"current_capacity":372,
	"minimum_capacity":145,"maximum_capacity":550,"warehouse_id":6,"product_type_id":34}`
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{ErrService: ErrRepo}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPost, "/sections", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// READ
// get all section in DB correctly
func TestSectionHandlerFindAll(t *testing.T) {
	//arrange
	objRes := responseSectionsTest{}
	serviceMock := section.ServiceMock{SectionsDummy: initData}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initData, objRes.Data)
}

// get all section in DB correctly but it has no data, send an informative message
func TestSectionHandlerFindAllAndNonData(t *testing.T) {
	//arrange
	expInfoMsg := "the data base has no records"
	objRes := map[string]string{}
	serviceMock := section.ServiceMock{}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expInfoMsg, objRes["data"])
}

// get all section but has problem with the repository
func TestSectionHandlerFindAllAndErrRepo(t *testing.T) {
	//arrange
	expErrMsg := "error repository down"
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{ErrService: ErrRepo}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// get a section by id correctly
func TestSectionHandlerFindByID(t *testing.T) {
	//arrange
	objRes := responseSectionTest{}
	serviceMock := section.ServiceMock{SectionDummy: initData[1]}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections/2", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initData[1], objRes.Data)
}

// get a section by id but id not exist
func TestSectionHandlerFindByIDNonExist(t *testing.T) {
	//arrange
	expErrMsg := "error section not found"
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{ErrService: s.ErrNotFound}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections/2", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// UPDATE
// update a section some or all attributes in the struct
func TestSectionHandlerUpdate(t *testing.T) {
	//arrange
	objRes := responseSectionTest{}
	body := `{"section_number":19,"current_temperature":24,"current_capacity":372}`
	serviceMock := section.ServiceMock{SectionDummy: initData[0]}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPatch, "/sections/2", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initData[0], objRes.Data)
}

// update a section some or all attributes in the struct
func TestSectionHandlerUpdateIDNonExist(t *testing.T) {
	//arrange
	expErrMsg := "error section not found"
	objRes := errorResponse{}
	body := `{"section_number":19,"current_temperature":24,"current_capacity":372}`
	serviceMock := section.ServiceMock{ErrService: s.ErrNotFound}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPatch, "/sections/2", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// update a section some or all attributes in the struct
func TestSectionHandlerUpdateAndSomeInvalidField(t *testing.T) {
	//arrange
	expErrMsg := "error json: cannot unmarshal string into Go struct field SectionPatch.section_number of type int"
	objRes := errorResponse{}
	body := `{"section_number":"otro","current_temperature":24,"current_capacity":372}`
	serviceMock := section.ServiceMock{}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPatch, "/sections/2", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// update a section but section number is invalid, it can´t be 0
func TestSectionHandlerUpdateAndInvalidSectionNumber(t *testing.T) {
	//arrange
	expErrMsg := "error not valid value: the section number could not be 0"
	body := `{"section_number":0,"current_temperature":24}`
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{ErrService: s.ErrInvalidValue{Msg: "the section number could not be 0"}}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodPatch, "/sections/2", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// DELETE
// Ddelete a section by id
func TestSectionHandlerDelete(t *testing.T) {
	//arrange
	serviceMock := section.ServiceMock{}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodDelete, "/sections/2", nil)
	//act
	router.ServeHTTP(rr, req)
	//assert
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body)
}

// delete a section but id not exist
func TestSectionHandlerDeleteIDNonExist(t *testing.T) {
	//arrange
	expErrMsg := "error section not found"
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{ErrService: s.ErrNotFound}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodDelete, "/sections/2", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}
