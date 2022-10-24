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
	data := []domain.Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	expectedData := []domain.Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	repository := mock.MockRepository{
		Data: data,
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
