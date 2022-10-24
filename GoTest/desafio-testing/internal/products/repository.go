package products

import "github.com/bootcamp-go/desafio-cierre-testing/internal/domain"

type Repository interface {
	GetAllBySeller(sellerID string) ([]domain.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	var prodList []domain.Product
	prodList = append(prodList, domain.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}
