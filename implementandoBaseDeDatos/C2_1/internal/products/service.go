package products

import "github.com/FPesenda/backpack-bcgow6-franco-pesenda/implementandoBaseDeDatos/C1_1/internal/domains"

type Service interface {
	Store(domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
	GetAll() ([]domains.Product, error)
	Delete(id int) error
	Update(domains.Product) error
	Exist(id int) bool
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

func (s *service) GetAll() (products []domains.Product, err error) {
	return s.repository.GetAll()
}

func (s *service) Delete(id int) (err error) {
	return s.repository.Delete(id)
}

func (s *service) Update(product domains.Product) (err error) {
	return s.repository.Update(product)
}

func (s *service) Exist(id int) bool {
	return s.repository.Exists(id)
}
