package mock

import "github.com/bootcamp-go/desafio-cierre-testing/internal/domain"

type MockRepository struct {
	Data  []domain.Product
	ErrId error
}

func (mock *MockRepository) GetAllBySeller(sellerID string) (products []domain.Product, err error) {
	if mock.ErrId != nil {
		err = mock.ErrId
		return
	}
	for _, v := range mock.Data {
		if v.SellerID == sellerID {
			products = append(products, v)
		}
	}
	return
}
