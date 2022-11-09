package warehouse

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	mocks "github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestGetAllHappy(t *testing.T) {
	//arrange
	expectedData := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "sarasa",
			Telephone:          "1142910831",
			WarehouseCode:      "test1",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		},
		{
			ID:                 2,
			Address:            "sarasa2",
			Telephone:          "114292132131",
			WarehouseCode:      "test2",
			MinimumCapacity:    10,
			MinimumTemperature: 10},
	}
	repository := mocks.RepositoryWarehouseMocks{
		Data: expectedData,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	result, err := service.GetAll(context.TODO())
	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData, result)
	assert.True(t, repository.GetAllFlag)

}

func TestGetAllSad(t *testing.T) {
	expectedError := errors.New("Repository error")
	repository := mocks.RepositoryWarehouseMocks{
		Data: nil,
		Err:  expectedError,
	}
	service := NewService(&repository)
	//act
	result, err := service.GetAll(context.TODO())
	//assert
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)
	assert.True(t, repository.GetAllFlag)

}

func TestGetByIdHappy(t *testing.T) {
	//arrange
	expectedData := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "sarasa",
			Telephone:          "1142910831",
			WarehouseCode:      "test1",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		},
		{
			ID:                 2,
			Address:            "sarasa2",
			Telephone:          "114292132131",
			WarehouseCode:      "test2",
			MinimumCapacity:    10,
			MinimumTemperature: 10},
	}
	repository := mocks.RepositoryWarehouseMocks{
		Data: expectedData,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	result, err := service.Get(context.TODO(), 1)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData[0], result)
	assert.True(t, repository.GetFlag)

}

func TestGetByIdSad(t *testing.T) {
	expectedError := errors.New("Repository error")
	expectedNilData := domain.Warehouse{}
	mockDb := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "sarasa",
			Telephone:          "1142910831",
			WarehouseCode:      "test1",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
		},
		{
			ID:                 2,
			Address:            "sarasa2",
			Telephone:          "114292132131",
			WarehouseCode:      "test2",
			MinimumCapacity:    10,
			MinimumTemperature: 10},
	}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  expectedError,
	}
	service := NewService(&repository)
	//act
	result, err := service.Get(context.TODO(), 3)
	//assert
	assert.Equal(t, expectedNilData, result)
	assert.Equal(t, expectedError, err)
	assert.True(t, repository.GetFlag)

}

func TestStoreHappy(t *testing.T) {
	//arrange
	warehouse := domain.Warehouse{
		ID:                 3,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	mockDb := []domain.Warehouse{}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	result, err := service.Save(context.TODO(), warehouse)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, 3, result)
	assert.True(t, repository.CreateFlag)

}

func TestStoreSad(t *testing.T) {
	//arrange
	warehouse := domain.Warehouse{
		ID:                 3,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}

	expectedError := errors.New("warehouseCode already used")
	mockDb := []domain.Warehouse{warehouse}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  expectedError,
	}
	service := NewService(&repository)
	//act
	result, err := service.Save(context.TODO(), warehouse)
	//assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, 0, result)
	assert.True(t, repository.CreateFlag)

}

func TestUpdateHappy(t *testing.T) {
	//arrange
	warehouse := domain.Warehouse{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	warehouseUpdated := []domain.Warehouse{{
		ID:                 1,
		Address:            "",
		Telephone:          "1142910831fterUpdate",
		WarehouseCode:      "sarasa1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	},
	}
	warehouseExpected := []domain.Warehouse{{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831fterUpdate",
		WarehouseCode:      "sarasa1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	},
	}
	mockDb := []domain.Warehouse{warehouse}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	err := service.Update(context.TODO(), warehouseUpdated[0])
	//assert
	assert.Nil(t, err)
	assert.Equal(t, warehouseExpected, repository.Data)
	assert.True(t, repository.UpdateFlag)

}

func TestUpdateSadWHinUse(t *testing.T) {
	//arrange
	warehouse := domain.Warehouse{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	warehouseTest := []domain.Warehouse{{
		ID:                 1,
		Address:            "",
		Telephone:          "1142910831fterUpdate",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	},
	}
	expectedError := errors.New("warehouseCode already in use")
	mockDb := []domain.Warehouse{warehouse}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	err := service.Update(context.TODO(), warehouseTest[0])
	//assert
	assert.Equal(t, expectedError, err)
	assert.True(t, repository.UpdateFlag)

}

func TestUpdateIDSad(t *testing.T) {
	//arrange
	warehouse := domain.Warehouse{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	warehouseTest := []domain.Warehouse{{
		ID:                 3,
		Address:            "",
		Telephone:          "1142910831fterUpdate",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	},
	}
	errId := fmt.Sprintf("warehouse %v id not found", warehouseTest[0].ID)
	expectedError := errors.New(errId)
	mockDb := []domain.Warehouse{warehouse}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	err := service.Update(context.TODO(), warehouseTest[0])
	//assert
	assert.Equal(t, expectedError, err)
	assert.True(t, repository.UpdateFlag)

}

func TestDeleteHappy(t *testing.T) {
	//arrange
	warehouse := domain.Warehouse{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	expectedData := []domain.Warehouse{}
	mockDb := []domain.Warehouse{warehouse}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  nil,
	}
	service := NewService(&repository)
	//act
	err := service.Delete(context.TODO(), 1)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData, repository.Data)
	assert.True(t, repository.DeleteFlag)

}
func TestDeleteSad(t *testing.T) {
	warehouse := domain.Warehouse{
		ID:                 1,
		Address:            "sarasa",
		Telephone:          "1142910831",
		WarehouseCode:      "test1",
		MinimumCapacity:    10,
		MinimumTemperature: 10,
	}
	expectedData := []domain.Warehouse{warehouse}
	mockDb := []domain.Warehouse{warehouse}
	repository := mocks.RepositoryWarehouseMocks{
		Data: mockDb,
		Err:  nil,
	}
	errId := fmt.Sprintf("warehouse %v id not found", 2)
	expectedError := errors.New(errId)
	service := NewService(&repository)
	//act
	err := service.Delete(context.TODO(), 2)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, expectedData, repository.Data)
	assert.True(t, repository.DeleteFlag)

}
