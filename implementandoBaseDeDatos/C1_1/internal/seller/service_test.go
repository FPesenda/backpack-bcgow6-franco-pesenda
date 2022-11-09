package seller

import (
	context "context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	mocks "github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/seller"
	"github.com/stretchr/testify/assert"
)

var initialData []domain.Seller = []domain.Seller{
	{
		ID:          1,
		CID:         33444555,
		CompanyName: "MELI",
		Address:     "sarasa 1234",
		Telephone:   "1122334455",
	},
	{
		ID:          2,
		CID:         44555666,
		CompanyName: "TESTING MELI",
		Address:     "test 1234",
		Telephone:   "1199887766",
	},
}

func TestFindAllOK(t *testing.T) {
	//arange
	repository := mocks.RepositorySellerMock{
		Data: initialData,
	}
	service := NewService(&repository)

	//act
	result, err := service.GetAll(context.TODO())
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, initialData, result)

}

func TestGetAllSad(t *testing.T) {
	//arrange
	expectedError := errors.New("error del repositorio")
	repository := mocks.RepositorySellerMock{
		Err: expectedError,
	}

	service := NewService(&repository)
	//act
	result, err := service.GetAll(context.TODO())
	//asert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}

func TestFindByIDExistent(t *testing.T) {
	//arrange

	expectedSeller := domain.Seller{
		ID:          1,
		CID:         33444555,
		CompanyName: "MELI",
		Address:     "sarasa 1234",
		Telephone:   "1122334455",
	}

	repository := mocks.RepositorySellerMock{
		Data: initialData,
	}
	service := NewService(&repository)
	var ctx context.Context

	//act
	result, err := service.Get(ctx, 1)

	//assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedSeller, result)

}
func TestFindByIDnonExistent(t *testing.T) {
	//arrange
	expectedError := errors.New("seller not found")

	repository := mocks.RepositorySellerMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	var ctx context.Context

	//act
	result, err := service.Get(ctx, 3)

	//assert
	assert.Empty(t, result)
	assert.NotNil(t, result)
	assert.Equal(t, expectedError, err)

}

func TestDeleteOK(t *testing.T) {
	//arange

	repository := mocks.RepositorySellerMock{
		Data: initialData,
	}
	service := NewService(&repository)

	//act
	err := service.Delete(context.TODO(), 1)

	//assert
	assert.Nil(t, err)
	//assert.NotNil(t, result)
	assert.Equal(t, 1, len(repository.Data))

}
func TestDeleteNonExistent(t *testing.T) {
	//arrange
	expectedError := errors.New("seller not found")

	repository := mocks.RepositorySellerMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	var ctx context.Context

	//act
	result := service.Delete(ctx, 3)

	//assert
	assert.NotNil(t, result)
	assert.Equal(t, expectedError, result)

}

func TestCreateSellerOK(t *testing.T) {
	//arrage
	expectedSeller := domain.Seller{
		ID:          1,
		CID:         33444555,
		CompanyName: "MELI",
		Address:     "sarasa 1234",
		Telephone:   "1122334455",
	}
	data, _ := json.Marshal(expectedSeller)
	fmt.Println(string(data))

	repository := mocks.RepositorySellerMock{}
	service := NewService(&repository)

	newSeller := domain.Seller{
		CID:         33444555,
		CompanyName: "MELI",
		Address:     "sarasa 1234",
		Telephone:   "1122334455",
	}
	//act
	result, err := service.Create(context.TODO(), newSeller)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedSeller, result)
}

func TestCreateSellerConflict(t *testing.T) {
	//arrage
	expectedSeller := domain.Seller{
		ID:          3,
		CID:         11221122,
		CompanyName: "MELI 123",
		Address:     "sarasa 2222",
		Telephone:   "22221111",
	}
	data, _ := json.Marshal(expectedSeller)
	fmt.Println(string(data))

	repository := mocks.RepositorySellerMock{
		Data: initialData,
	}
	service := NewService(&repository)

	newSeller := domain.Seller{
		CID:         11221122,
		CompanyName: "MELI 123",
		Address:     "sarasa 2222",
		Telephone:   "22221111",
	}
	//act
	result, err := service.Create(context.TODO(), newSeller)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedSeller, result)
}
func TestUpdateOK(t *testing.T) {
	//arrage
	expectedSeller := domain.Seller{
		ID:          2,
		CID:         908090877,
		CompanyName: "UPDATED MELI",
		Address:     "test 1234",
		Telephone:   "1199887766",
	}
	repository := mocks.RepositorySellerMock{
		Data: initialData,
	}
	service := NewService(&repository)

	updated := domain.Seller{
		ID:          2,
		CID:         908090877,
		CompanyName: "UPDATED MELI",
		Address:     "test 1234",
		Telephone:   "1199887766",
	}
	//act
	result, err := service.Update(context.TODO(), updated)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedSeller, result)
}
func TestUpdateNonExistent(t *testing.T) {
	//Arange
	expectedError := errors.New("sql: no rows in result set")
	repository := mocks.RepositorySellerMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	//Act

	result, err := service.Update(context.TODO(), domain.Seller{})
	//Assert
	assert.Equal(t, domain.Seller{}, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}
