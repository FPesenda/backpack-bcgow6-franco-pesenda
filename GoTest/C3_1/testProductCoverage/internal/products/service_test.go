package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	data       []Products
	readIsUsed bool
	errRead    error
	errWrite   error
}

func (d *MockStorage) Read(data interface{}) error {
	d.readIsUsed = true
	products := data.(*[]Products)
	*products = d.data
	return nil
}
func (d *MockStorage) Write(data interface{}) error {
	products := data.([]Products)
	d.data = products
	return nil
}

func TestUpdate(t *testing.T) {
	//ARANGE
	data := []Products{
		{
			Id:    1,
			Name:  "Mouse1",
			Color: "Black1",
			Code:  "Ms11",
			Price: 101.5,
		},
	}
	mockStorage := MockStorage{
		data:       data,
		readIsUsed: false,
		errRead:    nil,
		errWrite:   nil,
	}
	expectedProduct := Products{
		Id:    1,
		Name:  "MouseUpdated",
		Color: "BlackUpdated",
		Code:  "MsUpdated",
		Price: 101.5,
	}
	//ACT
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.UpdateByID(
		expectedProduct.Id,
		expectedProduct.Name,
		expectedProduct.Color,
		expectedProduct.Code,
		expectedProduct.Price,
	)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, mockStorage.readIsUsed)
	assert.Equal(t, expectedProduct, result, "deben ser iguales")
}

func TestHappyDelete(t *testing.T) {
	//ARANGE
	data := []Products{
		{
			Id:    1,
			Name:  "Mouse1",
			Color: "Black1",
			Code:  "Ms11",
			Price: 101.5,
		},
	}
	mockStorage := MockStorage{
		data:       data,
		readIsUsed: false,
		errRead:    nil,
		errWrite:   nil,
	}
	//ACT
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	err := service.Delete(1)
	assert.Nil(t, err)
}

func TestSadDelete(t *testing.T) {
	//ARANGE
	data := []Products{}
	mockStorage := MockStorage{
		data:       data,
		readIsUsed: false,
		errRead:    nil,
		errWrite:   nil,
	}
	idToDelete := 1
	errorExpected := fmt.Errorf("No se encuentra el elemento que se quiere borrar con id %d", idToDelete)
	//ACT
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	err := service.Delete(idToDelete)
	//ASSERT
	assert.NotNil(t, err)
	assert.Equal(t, errorExpected, err)
}
