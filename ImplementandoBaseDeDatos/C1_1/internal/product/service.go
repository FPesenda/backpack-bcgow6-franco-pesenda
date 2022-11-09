package product

import (
	"context"
	"errors"
)

// Errors
var (
	ErrNotFound     = errors.New("product not found")
	ErrProductExist = errors.New("the resource alredy exist in the database")
	ErrSaveFalied   = errors.New("fail in create new product")
)

type Service interface {
	Save(context.Context, domain.Product) (int, error)
	GetAll(context.Context) ([]domain.Product, error)
	Get(context.Context, int) (domain.Product, error)
	Update(context.Context, domain.Product) (domain.Product, error)
	Delete(context.Context, int) error
	Exist(context.Context, string) bool
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service *service) Save(ctx context.Context, p domain.Product) (saved int, err error) {
	if service.repository.Exists(ctx, p.ProductCode) {
		err = ErrProductExist
		return
	}
	saved, errSave := service.repository.Save(ctx, p)
	if errSave != nil {
		err = ErrSaveFalied
		return
	}
	return
}

func (service *service) GetAll(ctx context.Context) (productList []domain.Product, err error) {
	productList, errGetAll := service.repository.GetAll(ctx)
	if errGetAll != nil {
		err = errGetAll
		return
	}
	return
}

func (service *service) Get(ctx context.Context, id int) (product domain.Product, err error) {
	product, errGet := service.repository.Get(ctx, id)
	if errGet != nil {
		err = ErrNotFound
		return
	}
	return
}

func (service *service) Update(ctx context.Context, product domain.Product) (updated domain.Product, err error) {
	errUpdate := service.repository.Update(ctx, product)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	updated = product
	return
}

func (service *service) Delete(ctx context.Context, id int) (err error) {
	_, err = service.repository.Get(ctx, id)
	if err != nil {
		return
	}
	err = service.repository.Delete(ctx, id)
	return
}

func (service *service) Exist(ctx context.Context, productCode string) bool {
	return service.repository.Exists(ctx, productCode)
}
