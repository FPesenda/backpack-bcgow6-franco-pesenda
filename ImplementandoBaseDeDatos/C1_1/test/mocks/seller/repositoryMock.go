package mocks

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
	ErrExists   = errors.New("seller number exists")
	//ErrNotValidValue = errors.New("Not valid value")
)

type RepositorySellerMock struct {
	Data      []domain.Seller
	Err       error
	ErrorRepo error
}

func (repository *RepositorySellerMock) GetAll(ctx context.Context) ([]domain.Seller, error) {
	if repository.Err != nil {
		return nil, repository.Err
	}
	return repository.Data, nil
}

func (repository *RepositorySellerMock) Get(ctx context.Context, id int) (domain.Seller, error) {
	for i := range repository.Data {
		if repository.Data[i].ID == id {
			return repository.Data[i], nil
		}
	}

	return domain.Seller{}, ErrNotFound
}

func (repository *RepositorySellerMock) Exists(ctx context.Context, cid int) bool {
	for _, seller := range repository.Data {
		if seller.CID == cid {
			return true
		}
	}
	return false
}
func (repository *RepositorySellerMock) Save(ctx context.Context, s domain.Seller) (int, error) {
	if repository.ErrorRepo != nil {
		return 0, repository.ErrorRepo
	}
	s.ID = repository.nextID()
	repository.Data = append(repository.Data, s)
	return s.ID, nil
}

func (repository *RepositorySellerMock) Update(ctx context.Context, s domain.Seller) error {
	if repository.ErrorRepo != nil {
		return repository.ErrorRepo
	}
	for i := range repository.Data {
		if repository.Data[i].ID == s.ID {
			repository.Data[i] = s
			return nil
		}
	}
	return sql.ErrNoRows
}

func (repository *RepositorySellerMock) Delete(ctx context.Context, id int) error {

	for i := range repository.Data {
		if repository.Data[i].ID == id {
			repository.Data = append(repository.Data[:i], repository.Data[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
func (repository *RepositorySellerMock) nextID() int {
	if repository.Data == nil || len(repository.Data) == 0 {
		return 1
	}
	return repository.Data[len(repository.Data)-1].ID + 1
}
