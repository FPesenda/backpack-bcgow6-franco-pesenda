package seller

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

// Errors
var (
	ErrNotFound      = errors.New("seller not found")
	ErrExists        = errors.New("CID exists")
	ErrNotValidValue = errors.New("Not valid value")
)

type Service interface {
	Get(ctx context.Context, id int) (domain.Seller, error)
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Create(ctx context.Context, s domain.Seller) (domain.Seller, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, seller domain.Seller) (domain.Seller, error)
	Exist(ctx context.Context, cid int) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	sellers, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return sellers, nil
}

func (s *service) Create(ctx context.Context, seller domain.Seller) (domain.Seller, error) {

	if s.repository.Exists(ctx, seller.CID) {
		return domain.Seller{}, errors.New("CID already exists")
	}

	id, err := s.repository.Save(ctx, seller)
	if err != nil {
		return domain.Seller{}, err
	}
	seller.ID = id
	return seller, nil

}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return err
}

func (s *service) Get(ctx context.Context, id int) (domain.Seller, error) {
	seller, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Seller{}, err
	}
	return seller, nil
}

func (s *service) Update(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	err := s.repository.Update(ctx, seller)
	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

func (service *service) Exist(ctx context.Context, cid int) bool {
	return service.repository.Exists(ctx, cid)
}
