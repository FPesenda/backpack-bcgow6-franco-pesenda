package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type ServiceProductMock struct {
	Data              []domain.Product
	Err               error
	ErrUpdate         error
	ErrSave           error
	ErrFindById       error
	ErrDelete         error
	GetByIdIdExecuted bool
	SaveIsExecuted    bool
	UpdateIsExecuted  bool
}

func (service *ServiceProductMock) Save(ctx context.Context, prod domain.Product) (int, error) {
	service.SaveIsExecuted = true
	if service.Err != nil {
		return 0, service.Err
	}
	if service.ErrSave != nil {
		return 0, service.ErrSave
	}
	prod.ID = 1
	service.Data = append(service.Data, prod)
	return prod.ID, nil
}
func (service *ServiceProductMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	if service.Err != nil {
		return []domain.Product{}, service.Err
	}
	return service.Data, nil
}
func (service *ServiceProductMock) Get(ctx context.Context, id int) (domain.Product, error) {
	service.GetByIdIdExecuted = true
	if service.Err != nil {
		return domain.Product{}, service.Err
	}
	if service.ErrFindById != nil {
		return domain.Product{}, service.ErrFindById
	}
	for _, v := range service.Data {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.Product{}, service.Err
}
func (service *ServiceProductMock) Update(ctx context.Context, p domain.Product) (domain.Product, error) {
	service.UpdateIsExecuted = true
	if service.ErrUpdate != nil {
		return domain.Product{}, service.ErrUpdate
	}
	for _, v := range service.Data {
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
			return v, nil
		}
	}
	return domain.Product{}, service.ErrUpdate
}
func (service *ServiceProductMock) Delete(ctx context.Context, id int) error {
	if service.Err != nil {
		return service.Err
	}
	if service.ErrDelete != nil {
		return service.ErrDelete
	}
	return nil
}
func (service *ServiceProductMock) Exist(ctx context.Context, productCode string) bool {
	if service.Err != nil {
		return false
	}
	for _, v := range service.Data {
		if v.ProductCode == productCode {
			return true
		}
	}
	return true
}
