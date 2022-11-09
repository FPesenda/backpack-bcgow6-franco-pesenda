package mocks

import (
	//"context"
	//"fmt"

	"context"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type ServiceMock struct {
	SellersDummy     []domain.Seller
	SellerDummy      domain.Seller
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

func (r *ServiceMock) Create(ctx context.Context, s domain.Seller) (domain.Seller, error) {
	r.CreateFlag = true
	if r.ErrExists != nil {
		r.CreateFlag = false
		return domain.Seller{}, r.ErrExists
	}
	if r.ErrRepo != nil {
		return domain.Seller{}, r.ErrRepo
	}
	//r.SellerDummy.ID = s.ID
	return r.SellerDummy, nil

}

func (r *ServiceMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	r.GetAllFlag = true
	if r.ErrRepo != nil {
		return nil, r.ErrRepo
	}
	return r.SellersDummy, nil
}

func (r *ServiceMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	r.GetFlag = true
	if r.ErrNotFound != nil {
		return domain.Seller{}, r.ErrNotFound
	}
	if r.ErrRepo != nil {
		return domain.Seller{}, r.ErrRepo
	}
	return r.SellerDummy, nil
}

func (r *ServiceMock) Update(ctx context.Context, s domain.Seller) (domain.Seller, error) {
	r.UpdateFlag = true
	if r.ErrNotFound != nil {
		return domain.Seller{}, r.ErrNotFound
	}
	if r.ErrNotValidValue != nil {
		return domain.Seller{}, r.ErrNotValidValue
	}
	if r.ErrRepo != nil {
		return domain.Seller{}, r.ErrRepo
	}
	return r.SellerDummy, nil
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
func (r *ServiceMock) Exist(ctx context.Context, cid int) bool {

	return r.ExistsFlag
}
