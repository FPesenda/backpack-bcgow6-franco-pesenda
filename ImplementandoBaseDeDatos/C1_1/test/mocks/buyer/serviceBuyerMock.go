package buyer

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type ServiceMock struct {
	BuyersDummy      []domain.Buyer
	BuyerDummy       domain.Buyer
	GetAllFlag       bool
	GetFlag          bool
	ExistsFlag       bool
	CreateFlag       bool
	UpdateFlag       bool
	DeleteFlag       bool
	ErrNotFound      error
	ErrExists        error
	ErrRepo          error
	ErrNotValidValue error
}

func (r *ServiceMock) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	r.GetAllFlag = true
	if r.ErrRepo != nil {
		return nil, r.ErrRepo
	}
	return r.BuyersDummy, nil
}
func (r *ServiceMock) Get(ctx context.Context, id int) (domain.Buyer, error) {
	r.GetFlag = true
	if r.ErrNotFound != nil {
		return domain.Buyer{}, r.ErrNotFound
	}
	if r.ErrRepo != nil {
		return domain.Buyer{}, r.ErrRepo
	}
	return r.BuyerDummy, nil
}
func (r *ServiceMock) Save(ctx context.Context, s domain.Buyer) (int, error) {
	r.CreateFlag = true
	if r.ErrExists != nil {
		return 0, r.ErrExists
	}
	if r.ErrRepo != nil {
		return 0, r.ErrRepo
	}
	return r.BuyerDummy.ID, nil
}
func (r *ServiceMock) Update(ctx context.Context, s domain.Buyer) error {
	r.UpdateFlag = true
	if r.ErrNotFound != nil {
		return r.ErrNotFound
	}
	if r.ErrExists != nil {
		return r.ErrExists
	}
	if r.ErrNotValidValue != nil {
		return r.ErrNotValidValue
	}
	if r.ErrRepo != nil {
		return r.ErrRepo
	}
	return nil
}
func (r *ServiceMock) Delete(ctx context.Context, id int) error {
	r.DeleteFlag = true
	if r.ErrNotFound != nil {
		return r.ErrNotFound
	}
	if r.ErrRepo != nil {
		return r.ErrRepo
	}
	return nil
}
