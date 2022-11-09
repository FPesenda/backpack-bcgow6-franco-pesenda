package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type ServiceMock struct {
	DBDummy          []domain.Warehouse
	DataDummy        domain.Warehouse
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

func (s *ServiceMock) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	s.GetFlag = true
	if s.ErrNotFound != nil {
		return domain.Warehouse{}, s.ErrNotFound
	}
	if s.ErrRepo != nil {
		return domain.Warehouse{}, s.ErrRepo

	}
	for _, wh := range s.DBDummy {
		if wh.ID == id {
			return wh, nil
		}
	}
	return domain.Warehouse{}, errors.New("id not found")
}

func (s *ServiceMock) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	s.GetAllFlag = true
	if s.ErrNotFound != nil {
		return []domain.Warehouse{}, s.ErrNotFound
	}
	if s.ErrRepo != nil {
		return []domain.Warehouse{}, s.ErrRepo

	}
	if len(s.DBDummy) == 0 {
		s.DBDummy = nil
		return s.DBDummy, nil
	}
	return s.DBDummy, nil
}

func (s *ServiceMock) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	s.CreateFlag = true
	if s.ErrNotFound != nil {
		return 0, s.ErrNotFound
	}
	if s.ErrRepo != nil {
		return 0, s.ErrRepo

	}
	w.ID = 3
	s.DBDummy = append(s.DBDummy, w)
	return w.ID, nil

}

func (s *ServiceMock) Update(ctx context.Context, w domain.Warehouse) error {
	s.UpdateFlag = true

	s.DBDummy[w.ID-1] = w

	return nil
}

func (s *ServiceMock) Delete(ctx context.Context, id int) error {
	s.DeleteFlag = true

	s.DBDummy = append(s.DBDummy[:id-1], s.DBDummy[id:]...)

	return nil
}

func (s *ServiceMock) Exists(ctx context.Context, warehouseCode string) bool {
	for _, wh := range s.DBDummy {
		if warehouseCode == wh.WarehouseCode {
			return true
		}
	}
	return false
}
