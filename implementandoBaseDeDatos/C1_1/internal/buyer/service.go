package buyer

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
	Save(ctx context.Context, b domain.Buyer) (int, error)
	Update(ctx context.Context, b domain.Buyer) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) Get(ctx context.Context, id int) (domain.Buyer, error) {
	buyers, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Buyer{}, err
	}
	return buyers, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(ctx)
	if buyers == nil {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return buyers, nil
}
func (obj *service) Save(ctx context.Context, b domain.Buyer) (int, error) {
	if obj.repository.Exists(ctx, b.CardNumberID) {
		return 0, errors.New("card number id already exists")
	}
	id, err := obj.repository.Save(ctx, b)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(ctx context.Context, b domain.Buyer) error {
	err := s.repository.Update(ctx, b)
	if err != nil {
		return err
	}

	return nil
}
