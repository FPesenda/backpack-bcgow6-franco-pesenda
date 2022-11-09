package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/buyer"
	"github.com/stretchr/testify/assert"
)

type responseBuyerTest struct {
	Data domain.Buyer `json:"data"`
}
type responseBuyersTest struct {
	Data []domain.Buyer `json:"data"`
}

// Errors
var (
	ErrBuyersNotFound      = errors.New("error buyer not found")
	ErrBuyersExists        = errors.New("buyer number exists")
	ErrBuyersNotValidValue = errors.New("not valid value")
)

var initBuyerData []domain.Buyer = []domain.Buyer{
	{
		ID:           1,
		CardNumberID: "12345678",
		FirstName:    "Milagros",
		LastName:     "Sassenus",
	},
	{
		ID:           2,
		CardNumberID: "12345228",
		FirstName:    "Micaela",
		LastName:     "Gomez",
	},
}

// CREATE

// create a Buyer correctly
func TestBuyerHandlerCreateOk(t *testing.T) {
	newBuyer := domain.Buyer{ID: 0, CardNumberID: "12344444", FirstName: "Carla", LastName: "Car"}
	body, _ := json.Marshal(newBuyer)
	//arrange
	objRes := responseBuyerTest{}
	serviceMock := buyer.ServiceMock{BuyersDummy: initBuyerData}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodPost, "/buyers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.CreateFlag)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, newBuyer, objRes.Data)
	// assert.Contains(t, serviceMock.BuyersDummy, objRes.Data)
}

// create a Buyer but but not all required fields are sent
func TestBuyerHandlerCreateWithoutSomeFieldRequired(t *testing.T) {
	//arrange
	expErrMsg := "error : Key: 'BuyerPost.CardNumberID' Error:Field validation for 'CardNumberID' failed on the 'required' tag"
	newBuyer := domain.Buyer{FirstName: "Carla", LastName: "Car"}
	body, _ := json.Marshal(newBuyer)
	objRes := errorResponse{}
	serviceMock := buyer.ServiceMock{}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodPost, "/buyers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.CreateFlag)
	// assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// create a Buyer but Buyer number exists in the data
func TestBuyerHandlerCreateAndBuyerNumberAlreadyExist(t *testing.T) {
	//arrange
	expErrMsg := "error : buyer number exists"
	newBuyer := domain.Buyer{ID: 1, CardNumberID: "12345678", FirstName: "Carla", LastName: "Car"}
	body, _ := json.Marshal(newBuyer)
	objRes := errorResponse{}
	serviceMock := buyer.ServiceMock{ErrExists: errors.New("error : buyer number exists")}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodPost, "/buyers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.CreateFlag)
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// READ
// get all Buyer in DB correctly
func TestBuyerHandlerFindAll(t *testing.T) {
	//arrange
	objRes := responseBuyersTest{}
	serviceMock := buyer.ServiceMock{BuyersDummy: initBuyerData}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodGet, "/buyers", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetAllFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initBuyerData, objRes.Data)
}

func TestBuyerHandlerFindAllAndNonData(t *testing.T) {
	//arrange
	objRes := map[string]string{}
	serviceMock := buyer.ServiceMock{}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodGet, "/buyers", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetAllFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "", objRes["data"])
}

// get all Buyer but has problem with the repository
func TestBuyerHandlerFindAllAndErrRepo(t *testing.T) {
	//arrange
	expErrMsg := "the data base has no records"
	objRes := errorResponse{}
	serviceMock := buyer.ServiceMock{ErrRepo: ErrRepo}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodGet, "/buyers", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetAllFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// get a Buyer by id correctly
func TestBuyerHandlerFindByID(t *testing.T) {
	//arrange
	objRes := responseBuyerTest{}
	serviceMock := buyer.ServiceMock{BuyerDummy: initBuyerData[1]}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodGet, "/buyers/1", nil)
	//act
	router.ServeHTTP(rr, req)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes.Data)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initBuyerData[1], objRes.Data)
}

// get a Buyer by id but id not exist
func TestBuyerHandlerFindByIDNonExist(t *testing.T) {
	//arrange
	expErrMsg := "error buyer not found"
	objRes := errorResponse{}
	serviceMock := buyer.ServiceMock{ErrNotFound: errors.New("error buyer not found")}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodGet, "/buyers/2", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// UPDATE
// update a Buyer some or all attributes in the struct
func TestBuyerHandlerUpdate(t *testing.T) {
	//arrange
	newBuyer := domain.Buyer{ID: 1, CardNumberID: "12344444", FirstName: "Carla", LastName: "Car"}
	body, _ := json.Marshal(newBuyer)
	objRes := responseBuyerTest{}
	serviceMock := buyer.ServiceMock{BuyerDummy: initBuyerData[0]}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodPatch, "/buyers/1", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotEqual(t, initBuyerData[0], objRes.Data)
}

// update a Buyer some or all attributes in the struct
func TestBuyerHandlerUpdateIDNonExist(t *testing.T) {
	//arrange
	utilized_id := "2"
	expErrMsg := "error : buyer not found " + utilized_id
	objRes := errorResponse{}
	newBuyer := domain.Buyer{CardNumberID: "12344444", FirstName: "Carla", LastName: "Car"}
	body, _ := json.Marshal(newBuyer)
	serviceMock := buyer.ServiceMock{ErrNotFound: errors.New("buyer not found")}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodPatch, "/buyers/"+utilized_id, []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	// assert.True(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// create a section some or all attributes in the struct
func TestBuyerHandlerCreateAndSomeInvalidField(t *testing.T) {
	//arrange
	expErrMsg := "error : Key: 'BuyerPost.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"
	objRes := errorResponse{}
	newBuyer := request.BuyerPost{CardNumberID: "12344444", LastName: "Car"}
	body, _ := json.Marshal(newBuyer)
	serviceMock := buyer.ServiceMock{}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodPost, "/buyers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	fmt.Println(rr.Body)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// DELETE
// Delete a Buyer by id
func TestBuyerHandlerDelete(t *testing.T) {
	//arrange
	serviceMock := buyer.ServiceMock{}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodDelete, "/buyers/2", nil)
	//act
	router.ServeHTTP(rr, req)
	//assert
	assert.True(t, serviceMock.DeleteFlag)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body)
	assert.NotContains(t, serviceMock.BuyersDummy, initBuyerData[1])
}

// delete a Buyer but id not exist
func TestBuyerHandlerDeleteIDNonExist(t *testing.T) {
	//arrange
	expErrMsg := "buyer not found"
	objRes := errorResponse{}
	serviceMock := buyer.ServiceMock{ErrNotFound: errors.New("buyer not found")}
	handler := NewBuyer(&serviceMock)
	router := CreateServer(handler, "/buyers")
	req, rr := CreateRequestTest(http.MethodDelete, "/buyers/2", nil)
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.DeleteFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, expErrMsg, objRes.Message)
}

// delete a Buyer but id not valid data
//
