package buyer

import (
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	mocks "github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/buyer"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createDummyContext() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c
}

func TestCreateOk(t *testing.T) {
	//arange
	newBuyer := domain.Buyer{
		ID:           0,
		CardNumberID: "88888888",
		FirstName:    "Carlos",
		LastName:     "Francella"}

	repositoryData :=
		[]domain.Buyer{
			{
				ID:           1,
				CardNumberID: "12345678",
				FirstName:    "Milagros",
				LastName:     "Sassenus",
			},
		}

	repository := mocks.RepositoryBuyerMock{
		Data: repositoryData,
	}
	service := NewService(&repository)
	//act
	newId, err := service.Save(createDummyContext(), newBuyer)
	//assert
	assert.Nil(t, err, err)
	assert.Equal(t, repository.Data[len(repository.Data)-1].ID, newId)
}

func TestCreateSameCardIDFail(t *testing.T) {
	//arange
	newBuyer := domain.Buyer{ID: 2,
		CardNumberID: "12345678",
		FirstName:    "Carlos",
		LastName:     "Francella"}

	repositoryData :=
		[]domain.Buyer{
			{
				ID:           1,
				CardNumberID: "12345678",
				FirstName:    "Milagros",
				LastName:     "Sassenus",
			},
		}

	repository := mocks.RepositoryBuyerMock{
		Data: repositoryData,
	}
	service := NewService(&repository)
	//act
	_, err := service.Save(createDummyContext(), newBuyer)
	//assert

	assert.NotNil(t, err, err)
}

func TestFindAll(t *testing.T) {
	//arange
	expectedData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}
	repository := mocks.RepositoryBuyerMock{
		Data: expectedData,
	}
	service := NewService(&repository)
	//act
	result, err := service.GetAll(createDummyContext())
	//assert
	assert.Nil(t, err)
	assert.NotEmptyf(t, result, "there are no buyers!")

}

func TestFindAllEmpty(t *testing.T) {
	//arange
	repository := mocks.RepositoryBuyerMock{
		Data: nil,
	}
	service := NewService(&repository)
	//act
	_, err := service.GetAll(createDummyContext())
	//assert
	assert.NotNil(t, err, err)

}

func TestGetNonExistentId(t *testing.T) {
	//arange
	expectedData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}

	repository := mocks.RepositoryBuyerMock{
		Data: expectedData,
	}
	service := NewService(&repository)
	//act
	_, err := service.Get(createDummyContext(), 67)
	//assert
	assert.NotNil(t, err, err)

}

func TestGetExistentId(t *testing.T) {
	//arange
	expectedData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}

	repository := mocks.RepositoryBuyerMock{
		Data: expectedData,
	}
	service := NewService(&repository)
	//act
	result, err := service.Get(createDummyContext(), 2)
	//assert
	assert.Nil(t, err, err)
	assert.Equal(t, expectedData[1], result)

}

func TestUpdateExistent(t *testing.T) {
	//arange
	repoData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}

	newBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "00000000",
		FirstName:    "Brocoli",
		LastName:     "Verde",
	}

	repository := mocks.RepositoryBuyerMock{
		Data: repoData,
	}
	service := NewService(&repository)
	//act
	err := service.Update(createDummyContext(), newBuyer)
	//assert
	assert.Nil(t, err, err)
	assert.Equal(t, newBuyer, repository.Data[newBuyer.ID-1])

}

func TestUpdateNonExistent(t *testing.T) {
	//arange
	repoData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}

	newBuyer := domain.Buyer{
		ID:           89,
		CardNumberID: "00000000",
		FirstName:    "Brocoli",
		LastName:     "Verde",
	}

	repository := mocks.RepositoryBuyerMock{
		Data: repoData,
	}
	service := NewService(&repository)
	//act
	err := service.Update(createDummyContext(), newBuyer)
	//assert
	assert.NotNil(t, err, err)
}

func TestDeleteNonExistent(t *testing.T) {
	//arange
	repoData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}

	repository := mocks.RepositoryBuyerMock{
		Data: repoData,
	}
	service := NewService(&repository)
	//act
	err := service.Delete(createDummyContext(), 10)
	//assert
	assert.NotNil(t, err, err)

}

func TestDeleteExistent(t *testing.T) {
	//arange
	repoData := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "12345678",
			FirstName:    "Milagros",
			LastName:     "Sassenus",
		},
		{
			ID:           2,
			CardNumberID: "12322278",
			FirstName:    "Carla",
			LastName:     "Rivero",
		},
	}

	repository := mocks.RepositoryBuyerMock{
		Data: repoData,
	}
	service := NewService(&repository)
	//act
	err := service.Delete(createDummyContext(), 2)
	//assert
	assert.Nil(t, err, err)
	assert.Equal(t, len(repository.Data), len(repoData)-1)
	assert.NotContains(t, repository.Data, repoData[len(repoData)-1])
	assert.Contains(t, repository.Data, repoData[0])

}
