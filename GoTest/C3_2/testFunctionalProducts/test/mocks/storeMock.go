package mocks

import (
	"fmt"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/products"
)

type MockStorage struct {
	Data     []products.Products
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]products.Products)
	*a = m.Data
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]products.Products)
	m.Data = append(m.Data, a[len(a)-1])
	return nil
}
