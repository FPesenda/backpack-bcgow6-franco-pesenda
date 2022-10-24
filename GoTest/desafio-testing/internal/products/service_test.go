package products

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-testing/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBySellerHappy(t *testing.T) {
	//ARANGE
	expectedData := []domain.Product{
		{
			ID:          "mock1",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
		{
			ID:          "mock2",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       113.55,
		},
	}
	repository := mock.MockRepository{
		Data: expectedData,
	}
	service := NewService(&repository)
	//ACT
	result, err := service.GetAllBySeller(expectedData[0].SellerID)
	//ASSERT
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedData, result)
}

func TestGetAllBySellerSad(t *testing.T) {
	//ARANGE
	sellerId := "sellerIDError"
	expectedError := errors.New(fmt.Sprint("error in repository sellerId:", sellerId))
	repository := mock.MockRepository{
		ErrId: expectedError,
	}
	service := NewService(&repository)
	//ACT
	result, err := service.GetAllBySeller(sellerId)
	//ASSERT
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}
