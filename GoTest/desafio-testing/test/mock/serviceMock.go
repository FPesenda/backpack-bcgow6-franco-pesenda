package mock

import "github.com/bootcamp-go/desafio-cierre-testing/internal/domain"

type ServiceMock struct {
	Data          []domain.Product
	ErrRepository error
}

func (mock *ServiceMock) GetAllBySeller(sellerID string) (products []domain.Product, err error) {
	if mock.ErrRepository != nil {
		err = mock.ErrRepository
		return
	}
	for _, v := range mock.Data {
		if v.SellerID == sellerID {
			products = append(products, v)
		}
	}
	return
}
