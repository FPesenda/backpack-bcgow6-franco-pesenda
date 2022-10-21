package mocks

import (
	"fmt"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/domain"
)

type MockStorage struct {
	Data     []domain.Product
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]domain.Product)
	*a = m.Data
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]domain.Product)
	m.Data = append(m.Data, a[len(a)-1])
	return nil
}
