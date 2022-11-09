package buyer

import (
	"context"
	"errors"
	"sort"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type RepositoryBuyerMock struct {
	Data []domain.Buyer
	Err  error
}

func (repository *RepositoryBuyerMock) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	if repository.Err != nil {
		return nil, repository.Err
	}
	return repository.Data, nil
}

func (repository *RepositoryBuyerMock) Get(ctx context.Context, id int) (domain.Buyer, error) {

	for _, buyer := range repository.Data {
		if buyer.ID == id {
			return buyer, nil
		}
	}
	return domain.Buyer{}, errors.New("error : id not found")
}

func (repository *RepositoryBuyerMock) Exists(ctx context.Context, cardNumberID string) bool {
	for _, buyer := range repository.Data {
		if buyer.CardNumberID == cardNumberID {
			return true
		}
	}
	return false
}
func (repository *RepositoryBuyerMock) Save(ctx context.Context, newBuyer domain.Buyer) (int, error) {
	newBuyer.ID = len(repository.Data) + 1
	repository.Data = append(repository.Data, newBuyer)
	return repository.Data[len(repository.Data)-1].ID, nil
}

func (repository *RepositoryBuyerMock) Update(ctx context.Context, updatedBuyer domain.Buyer) error {
	if (len(repository.Data) < updatedBuyer.ID) || updatedBuyer.ID < 1 {
		return errors.New("error : invalid id")
	}
	repository.Data[updatedBuyer.ID-1] = updatedBuyer
	return nil
}

func (repository *RepositoryBuyerMock) Delete(ctx context.Context, id int) error {
	if (len(repository.Data) < id) || id < 1 {
		return errors.New("error : invalid id")
	}
	repository.Data[id-1] = repository.Data[len(repository.Data)-1]
	repository.Data = repository.Data[:len(repository.Data)-1]
	sort.Slice(repository.Data, func(p, q int) bool {
		return repository.Data[p].ID < repository.Data[q].ID
	})
	return nil
}
