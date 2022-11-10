package products

import "github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/domains"

type Service interface {
	Store(domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(p domains.Product) (int, error) {
	return s.repository.Store(p)
}

func (s *service) GetByName(name string) (domains.Product, error) {
	return s.repository.GetByName(name)
}
