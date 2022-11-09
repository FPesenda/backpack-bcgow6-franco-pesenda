package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("warehouse not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w domain.Warehouse) (int, error)
	Update(ctx context.Context, w domain.Warehouse) error
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	warehouse, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return warehouse, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	warehouseList, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return warehouseList, nil
}

func (s *service) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	id, err := s.repository.Save(ctx, w)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) Update(ctx context.Context, w domain.Warehouse) error {
	err := s.repository.Update(ctx, w)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Exists(ctx context.Context, warehouseCode string) bool {
	return s.repository.Exists(ctx, warehouseCode)
}
