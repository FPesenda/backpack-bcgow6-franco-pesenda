package mocks

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type RepositoryProductMock struct {
	Data         []domain.Product
	DeleteIsUsed bool
	GetIsUsed    bool
	ErrDelete    error
	ErrNotFound  error
	Err          error
}

func (repository *RepositoryProductMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	if repository.Err != nil {
		return nil, repository.Err
	}
	return repository.Data, nil
}

func (repository *RepositoryProductMock) Get(ctx context.Context, id int) (domain.Product, error) {
	repository.GetIsUsed = true
	if repository.Err != nil {
		return domain.Product{}, repository.Err
	}
	for _, v := range repository.Data {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (repository *RepositoryProductMock) Exists(ctx context.Context, productCode string) bool {
	if repository.Err != nil {
		return false
	}
	for _, v := range repository.Data {
		if v.ProductCode == productCode {
			return true
		}
	}
	return false
}

func (repository *RepositoryProductMock) Save(ctx context.Context, p domain.Product) (int, error) {
	repository.DeleteIsUsed = true
	if repository.Err != nil {
		return 0, repository.Err
	}
	idNew := 0
	repository.Data = append(repository.Data, p)
	for _, v := range repository.Data {
		if v.ProductCode == p.ProductCode {
			idNew = v.ID
		}
	}
	return idNew, nil
}

func (repository *RepositoryProductMock) Update(ctx context.Context, p domain.Product) error {
	if repository.Err != nil {
		return repository.Err
	}
	if repository.ErrNotFound != nil {
		return repository.ErrNotFound
	}
	for _, v := range repository.Data {
		if v.ID == p.ID {
			v.Description = p.Description
			v.ExpirationRate = p.ExpirationRate
			v.FreezingRate = p.FreezingRate
			v.Height = p.Height
			v.Length = p.Length
			v.Netweight = p.Netweight
			v.ProductCode = p.ProductCode
			v.RecomFreezTemp = p.RecomFreezTemp
			v.Width = p.Width
			v.ProductTypeID = p.ProductTypeID
			v.SellerID = p.SellerID
			return nil
		}
	}
	return repository.ErrNotFound
}

func (repository *RepositoryProductMock) Delete(ctx context.Context, id int) error {
	repository.DeleteIsUsed = true
	if repository.ErrDelete != nil {
		return repository.ErrDelete
	}
	if repository.ErrNotFound != nil {
		return repository.ErrNotFound
	}
	for i := range repository.Data {
		if repository.Data[i].ID == id {
			repository.Data = append(repository.Data[:i], repository.Data[i+1:]...)
			return nil
		}
	}
	return repository.Err
}
