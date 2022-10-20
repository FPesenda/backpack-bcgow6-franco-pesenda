package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyDB struct{}

func (d DummyDB) Read(data interface{}) error {
	products := data.(*[]Products)
	*products = []Products{
		{
			Id:    1,
			Name:  "Mouse2",
			Color: "Black2",
			Code:  "Ms12",
			Price: 101.5,
		},
		{
			Id:    2,
			Name:  "Mouse3",
			Color: "Black3",
			Code:  "Ms13",
			Price: 101.5,
		},
	}
	return nil
}
func (d DummyDB) Write(data interface{}) error {
	return nil
}

type MockDB struct {
	readUSed bool
	data     []Products
}

func (d *MockDB) Read(data interface{}) error {
	d.readUSed = true
	products := data.(*[]Products)
	*products = d.data
	return nil
}
func (d *MockDB) Write(data interface{}) error {
	products := data.([]Products)
	d.data = products
	return nil
}

func TestGetAll(t *testing.T) {
	dumy := DummyDB{}
	repository := NewRepository(dumy)
	products, _ := repository.GetAll()
	productsExpected := []Products{
		{
			Id:    1,
			Name:  "Mouse2",
			Color: "Black2",
			Code:  "Ms12",
			Price: 101.5,
		},
		{
			Id:    2,
			Name:  "Mouse3",
			Color: "Black3",
			Code:  "Ms13",
			Price: 101.5,
		},
	}
	assert.Equal(t, productsExpected, products, "Deben ser iguales")
}

func TestUpdateName(t *testing.T) {
	data := []Products{
		{
			Id:    1,
			Name:  "Mouse2",
			Color: "Black2",
			Code:  "Ms12",
			Price: 101.5,
		},
	}
	NewName := "MOUSE RAZER"
	dataExpected := Products{
		Id:    1,
		Name:  NewName,
		Color: "Black2",
		Code:  "Ms12",
		Price: 101.5,
	}

	mock := MockDB{
		readUSed: false,
		data:     data,
	}
	repository := NewRepository(&mock)
	updated, err := repository.ChangeName(1, NewName)
	assert.Nil(t, err)
	assert.True(t, mock.readUSed)
	assert.Equal(t, dataExpected, updated)
}
