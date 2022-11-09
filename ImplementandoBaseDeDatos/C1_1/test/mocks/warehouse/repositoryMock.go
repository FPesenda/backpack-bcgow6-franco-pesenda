package warehouse

import (
	"context"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type RepositoryWarehouseMocks struct {
	Data       []domain.Warehouse
	Err        error
	GetAllFlag bool
	GetFlag    bool
	ExistsFlag bool
	CreateFlag bool
	UpdateFlag bool
	DeleteFlag bool
}

func (repository *RepositoryWarehouseMocks) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	repository.GetAllFlag = true
	if repository.Err != nil {
		return nil, repository.Err
	}
	return repository.Data, nil
}
func (repository *RepositoryWarehouseMocks) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	repository.GetFlag = true
	for _, wh := range repository.Data {
		if wh.ID == id {
			return repository.Data[id-1], nil
		}
	}
	return domain.Warehouse{}, repository.Err
}
func (repository *RepositoryWarehouseMocks) Exists(ctx context.Context, warehouseCode string) bool {
	return false
}
func (repository *RepositoryWarehouseMocks) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	repository.CreateFlag = true
	for _, wh := range repository.Data {
		if w.WarehouseCode == wh.WarehouseCode {
			return 0, repository.Err
		}
	}
	repository.Data = append(repository.Data, w)
	return w.ID, nil
}
func (repository *RepositoryWarehouseMocks) Update(ctx context.Context, w domain.Warehouse) error {
	repository.UpdateFlag = true
	for i := range repository.Data {
		if w.ID == repository.Data[i].ID {
			if w.WarehouseCode == repository.Data[i].WarehouseCode {
				return errors.New("warehouseCode already in use")
			}
			if w.WarehouseCode != "" {
				repository.Data[i].WarehouseCode = w.WarehouseCode
			}
			if w.Address != "" {
				repository.Data[i].Address = w.Address
			}
			if w.Telephone != "" {
				repository.Data[i].Telephone = w.Telephone
			}
			if w.MinimumCapacity < 0 {
				return repository.Err
			}
			if w.MinimumCapacity != 0 {
				repository.Data[i].MinimumCapacity = w.MinimumCapacity
			}
			if w.MinimumTemperature != 0 {
				repository.Data[i].MinimumTemperature = w.MinimumTemperature
			}

		} else {
			err := fmt.Sprintf("warehouse %v id not found", w.ID)
			return errors.New(err)
		}
	}
	return nil
}

func (repo *RepositoryWarehouseMocks) Delete(ctx context.Context, id int) error {
	repo.DeleteFlag = true
	for i := range repo.Data {
		if repo.Data[i].ID == id {
			repo.Data = append(repo.Data[:i], repo.Data[i+1:]...)
			return nil
		} else {
			err := fmt.Sprintf("warehouse %v id not found", id)
			return errors.New(err)
		}

	}
	return nil
}
