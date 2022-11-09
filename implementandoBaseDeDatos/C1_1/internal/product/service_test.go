package product

import (
	"context"
	"errors"
	"testing"
)

var ctx context.Context = context.TODO()

func TestGetAllHappy(t *testing.T) {
	//arange
	expectedData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "Descripciñon 2",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	repository := mocks.RepositoryProductMock{
		Data: expectedData,
	}
	service := NewService(&repository)
	//act
	result, err := service.GetAll(ctx)
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedData, result)
}

func TestGetAllSad(t *testing.T) {
	//arange
	expectedError := errors.New("error del repositorio")
	repository := mocks.RepositoryProductMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	//act
	result, err := service.GetAll(ctx)
	//asert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func TestGetByIdHappy(t *testing.T) {
	//arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "Descripciñon 2",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := domain.Product{
		ID:             2,
		Description:    "Descripciñon 2",
		ExpirationRate: 10,
		FreezingRate:   20,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	idToSerch := 2
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	service := NewService(&repository)
	//act
	result, err := service.Get(ctx, idToSerch)
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedData, result)
}

// PREGUNTAR SI ESTE TEST TIENE VALOR
func TestGetByIdSad(t *testing.T) {
	//Arange
	expectedError := errors.New("product not found")
	repository := mocks.RepositoryProductMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	//Act
	result, err := service.Get(ctx, 0)
	//Assert
	assert.Equal(t, result, domain.Product{})
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func TestGetByIdSadIdNonExistent(t *testing.T) {
	//arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := errors.New("product not found")
	idToSerch := 2
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	service := NewService(&repository)
	//act
	result, err := service.Get(ctx, idToSerch)
	//assert
	assert.NotNil(t, err)
	assert.Empty(t, result)
	assert.EqualError(t, expectedData, err.Error())
}

func TestSaveHappy(t *testing.T) {
	//arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "Descripciñon 2",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE2",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	productToAdd := domain.Product{
		ID:             2,
		Description:    "Descripciñon 2",
		ExpirationRate: 10,
		FreezingRate:   20,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE2",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	service := NewService(&repository)
	//act
	result, err := service.Save(ctx, productToAdd)
	//assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result, productToAdd.ID)
	assert.Equal(t, expectedData, repository.Data)
}

func TestSaveSad(t *testing.T) {
	//arange
	expectedError := errors.New("fail in create new product")
	repository := mocks.RepositoryProductMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	//act
	result, err := service.Save(ctx, domain.Product{})
	//assert
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func TestSaveSadExist(t *testing.T) {
	//arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	productToAdd := domain.Product{
		ID:             1,
		Description:    "Descripciñon 1",
		ExpirationRate: 10,
		FreezingRate:   20,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	service := NewService(&repository)
	//act
	result, err := service.Save(ctx, productToAdd)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, result, 0)
	assert.Equal(t, expectedData, repository.Data)
}

func TestUpdateHappy(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedProduct := domain.Product{
		ID:             1,
		Description:    "Descripciñon Updated",
		ExpirationRate: 99,
		FreezingRate:   20,
		Height:         99.9,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODEUpdated",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	service := NewService(&repository)
	//Act
	result, err := service.Update(ctx, expectedProduct)
	//Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedProduct, result)
}

func TestUpdateSad(t *testing.T) {
	//Arange
	expectedError := errors.New("fail in update product")
	repository := mocks.RepositoryProductMock{
		Err: expectedError,
	}
	service := NewService(&repository)
	//Act
	result, err := service.Update(ctx, domain.Product{})
	//Assert
	assert.Equal(t, domain.Product{}, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func TestUpdateSadIdNonExist(t *testing.T) {
	//Arange
	expectedError := errors.New("fail in update product")
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	productToAdd := domain.Product{
		ID:             10,
		Description:    "Descripciñon 1",
		ExpirationRate: 10,
		FreezingRate:   20,
		Height:         10.1,
		Length:         20.1,
		Netweight:      30.1,
		ProductCode:    "CODE1",
		RecomFreezTemp: 40.1,
		Width:          50.1,
		ProductTypeID:  1,
		SellerID:       1,
	}
	repository := mocks.RepositoryProductMock{
		Data:        initialData,
		ErrNotFound: expectedError,
	}
	service := NewService(&repository)
	//Act
	result, err := service.Update(ctx, productToAdd)
	//Assert
	assert.Equal(t, domain.Product{}, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func TestDeleteHappy(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
		{
			ID:             2,
			Description:    "Descripciñon 2",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	expectedData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	idToDelete := 2
	service := NewService(&repository)
	//Act
	err := service.Delete(ctx, idToDelete)
	//Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedData, repository.Data)
}

func TestDeleteSad(t *testing.T) {
	//Arange
	expectedError := errors.New("fail delete")
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	repository := mocks.RepositoryProductMock{
		Data:         initialData,
		ErrDelete:    expectedError,
		DeleteIsUsed: false,
		GetIsUsed:    false,
	}
	service := NewService(&repository)
	//Act
	err := service.Delete(ctx, 1)
	//Assert
	assert.True(t, repository.GetIsUsed)
	assert.True(t, repository.DeleteIsUsed)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())
}

func TestDeleteSadIdNonExistent(t *testing.T) {
	//Arange
	expectedError := ErrNotFound
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	idToDelete := 10
	repository := mocks.RepositoryProductMock{
		Data:        initialData,
		ErrNotFound: expectedError,
	}
	service := NewService(&repository)
	//Act
	err := service.Delete(ctx, idToDelete)
	//Assert
	assert.NotNil(t, err)
	assert.EqualError(t, expectedError, err.Error())
}

func TestExistHappy(t *testing.T) {
	//Arange
	initialData := []domain.Product{
		{
			ID:             1,
			Description:    "Descripciñon 1",
			ExpirationRate: 10,
			FreezingRate:   20,
			Height:         10.1,
			Length:         20.1,
			Netweight:      30.1,
			ProductCode:    "CODE1",
			RecomFreezTemp: 40.1,
			Width:          50.1,
			ProductTypeID:  1,
			SellerID:       1,
		},
	}
	codeToSerch := "CODE1"
	repository := mocks.RepositoryProductMock{
		Data: initialData,
	}
	service := NewService(&repository)
	//Act
	exs := service.Exist(ctx, codeToSerch)
	//Assert
	assert.True(t, exs)
}
