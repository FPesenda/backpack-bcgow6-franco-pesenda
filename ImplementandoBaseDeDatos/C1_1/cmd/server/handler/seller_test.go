package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	mocks "github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/seller"

	"github.com/stretchr/testify/assert"
)

type responseSellerTest struct {
	Data domain.Seller `json:"data"`
}
type responseSellersTest struct {
	Data []domain.Seller `json:"data"`
}
type errorResponseSeller struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

var ErrRepoSeller error = errors.New("repository down")

var initialData []domain.Seller = []domain.Seller{
	{
		ID:          1,
		CID:         33444555,
		CompanyName: "MELI",
		Address:     "sarasa 1234",
		Telephone:   "1122334455",
	},
	{
		ID:          2,
		CID:         44555666,
		CompanyName: "TESTING MELI",
		Address:     "test 1234",
		Telephone:   "1199887766",
	},
}

func TestSellerHandlerCreateOk(t *testing.T) {
	//arrange
	body := `{"cid": 33444555,"company_name": "MELI","address": "sarasa 1234", "telephone": "122334455"}`
	objRes := responseSellerTest{}
	serviceMock := mocks.ServiceMock{SellerDummy: initialData[0]}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPost, "/sellers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.CreateFlag)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, initialData[0], objRes.Data)

}

func TestSellerHandlerCreateConflict(t *testing.T) {
	errMsg := "seller with CID 44555666 already exists"

	body := `{"cid": 44555666,"company_name": "MELI","address": "sarasa 1234", "telephone": "122334455"}`
	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ErrExists: mocks.ErrExists}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPost, "/sellers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.CreateFlag)
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}

func TestSellerHandlerCreateBadRequest(t *testing.T) {
	errMsg := "error : not a json"

	body := `{"cid": "44555666","company_name": "MELI","address": "sarasa 1234", "telephone": "122334455"}`
	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ErrNotValidValue: mocks.ErrNotFound}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPost, "/sellers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.CreateFlag)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}
func TestSellerHandlerCreateFail(t *testing.T) {
	errMsg := "error : the CID is required"

	body := `{"company_name": "MELI","address": "sarasa 1234", "telephone": "122334455"}`
	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ErrExists: mocks.ErrExists}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPost, "/sellers", []byte(body))
	//act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	//assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.CreateFlag)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}
func TestFindAllOK(t *testing.T) {
	// arrange
	objRes := responseSellersTest{}
	serviceMock := mocks.ServiceMock{SellersDummy: initialData}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodGet, "/sellers", nil)
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetAllFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initialData, objRes.Data)

}
func TestFindAByIDexistent(t *testing.T) {
	// arrange

	objRes := responseSellerTest{}
	serviceMock := mocks.ServiceMock{SellerDummy: initialData[1]}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodGet, "/sellers/1", nil)
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initialData[1], objRes.Data)

}
func TestFindAByIDnonexistent(t *testing.T) {
	// arrange
	errMsg := "seller with 3 does not exist "

	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ErrNotFound: mocks.ErrNotFound}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodGet, "/sellers/3", nil)
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.GetFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}
func TestUpdateOK(t *testing.T) {
	// arrange
	body := `{"cid": 10,"company_name": "COMPANIA ANONIMA"}`

	objRes := responseSellerTest{}
	serviceMock := mocks.ServiceMock{SellerDummy: initialData[1]}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPatch, "/sellers/1", []byte(body))
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initialData[1], objRes.Data)

}
func TestUpdatePutOK(t *testing.T) {
	// arrange
	body := `{"cid": 1000,"company_name": "PUT","address": "sarasa 1234", "telephone": "122334455"}`

	objRes := responseSellerTest{}
	serviceMock := mocks.ServiceMock{SellerDummy: initialData[1]}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPatch, "/sellers/1", []byte(body))
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, initialData[1], objRes.Data)

}
func TestUpdateConflict(t *testing.T) {
	// arrange
	errMsg := "seller with CID 33444555 already exists "

	body := `{"cid": 33444555,"company_name": "PUT","address": "sarasa 1234", "telephone": "122334455"}`

	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ExistsFlag: true}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPatch, "/sellers/1", []byte(body))
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}
func TestUpdateNonExistent(t *testing.T) {
	// arrange
	errMsg := "seller with 5 does not exist "
	body := `{"cid": 10,"company_name": "COMPANIA ANONIMA"}`

	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ErrNotFound: mocks.ErrNotFound}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodPatch, "/sellers/5", []byte(body))
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.UpdateFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}

func TestDeleteExistentOK(t *testing.T) {
	//arrange
	serviceMock := mocks.ServiceMock{}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodDelete, "/sellers/1", nil)
	//act
	router.ServeHTTP(rr, req)
	//assert
	assert.True(t, serviceMock.DeleteFlag)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body)

}

func TestDeleteNonexistent(t *testing.T) {
	// arrange
	errMsg := "seller with 3 does not exist "

	objRes := errorResponseSeller{}
	serviceMock := mocks.ServiceMock{ErrNotFound: mocks.ErrNotFound}
	handler := NewSeller(&serviceMock)
	router := CreateServer(handler, "/sellers")
	req, rr := CreateRequestTest(http.MethodDelete, "/sellers/3", nil)
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.True(t, serviceMock.DeleteFlag)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, errMsg, objRes.Message)

}
