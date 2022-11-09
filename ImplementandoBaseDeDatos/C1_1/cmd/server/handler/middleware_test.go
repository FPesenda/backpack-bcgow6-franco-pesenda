package handler

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/section"
	"github.com/stretchr/testify/assert"
)

// middleware token validation, token dosen´t mached
func TestMiddlewareInvalidToken(t *testing.T) {
	// arrange
	expInfoMsg := "error does not have permissions to perform the requested request, invalid token: 123"
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections", nil)
	req.Header.Set("token", "123")
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, expInfoMsg, objRes.Message)
}

// middleware id is invalid isn´t a integer
func TestMiddlewareInvalidId(t *testing.T) {
	// arrange
	expInfoMsg := "error invalid ID: aaa"
	objRes := errorResponse{}
	serviceMock := section.ServiceMock{}
	handler := NewSection(&serviceMock)
	router := CreateServer(handler, endPoint)
	req, rr := CreateRequestTest(http.MethodGet, "/sections/aaa", nil)
	// act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// assert
	assert.Nil(t, err)
	assert.False(t, serviceMock.ServiceMethodFlag)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expInfoMsg, objRes.Message)
}
