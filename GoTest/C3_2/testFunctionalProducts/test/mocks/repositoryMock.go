package mocks

import (
	"fmt"

	"github.com/FPesenda/backpack-bcgow6-franco-pesenda/GoTest/C3_2/testFunctionalProducts/internal/domain"
)

type MockRepository struct {
	Data  []domain.Product
	Error string
}

func (m *MockRepository) GetAll() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.Data, nil
}

func (m *MockRepository) Store(id int, name, producType string, count int, price float64) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}
	p := domain.Product{
		ID:    id,
		Name:  name,
		Type:  producType,
		Count: count,
		Price: price,
	}
	return p, nil
}

func (m *MockRepository) LastID() (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}
	return m.Data[len(m.Data)-1].ID, nil
}
