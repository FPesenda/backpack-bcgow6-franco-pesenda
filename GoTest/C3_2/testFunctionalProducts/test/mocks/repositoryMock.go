package mocks

import (
	"fmt"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/products"
)

type MockRepository struct {
	Data  []products.Products
	Error string
}

func (m *MockRepository) GetAll() ([]products.Products, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.Data, nil
}

func (m *MockRepository) Store(id int, name, color string, code string, price float64) (products.Products, error) {
	if m.Error != "" {
		return products.Products{}, fmt.Errorf(m.Error)
	}
	p := products.Products{
		Id:    id,
		Name:  name,
		Color: color,
		Code:  code,
		Price: price,
	}
	return p, nil
}
func (m *MockRepository) LastID() (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}
	return m.Data[len(m.Data)-1].Id, nil
}
