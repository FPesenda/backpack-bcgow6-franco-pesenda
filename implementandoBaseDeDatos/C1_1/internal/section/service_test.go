package section

import (
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/test/mocks/section"
	"github.com/stretchr/testify/assert"
)

var initData []domain.Section = []domain.Section{
	{
		ID:                 1,
		SectionNumber:      13,
		CurrentTemperature: 27,
		MinimumTemperature: 10,
		CurrentCapacity:    143,
		MinimumCapacity:    100,
		MaximumCapacity:    500,
		WarehouseID:        3,
		ProductTypeID:      2,
	},
	{
		ID:                 2,
		SectionNumber:      14,
		CurrentTemperature: 12,
		MinimumTemperature: 8,
		CurrentCapacity:    25,
		MinimumCapacity:    1,
		MaximumCapacity:    250,
		WarehouseID:        7,
		ProductTypeID:      16,
	},
	{
		ID:                 3,
		SectionNumber:      15,
		CurrentTemperature: 13,
		MinimumTemperature: 5,
		CurrentCapacity:    14,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseID:        5,
		ProductTypeID:      12,
	},
}
var ErrRepo error = errors.New("error in repository")

func initialDataCopy() []domain.Section {
	data := make([]domain.Section, len(initData))
	copy(data, initData)
	return data
}

// CREATE
// test to create section when data initial is empty
func TestSectionServiceCreateOkEmptyData(t *testing.T) {
	//arrange
	expSection := domain.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 15,
		MinimumTemperature: 8,
		CurrentCapacity:    14,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseID:        5,
		ProductTypeID:      12,
	}
	repoMock := section.RepositoryMock{}
	service := NewService(&repoMock)
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      1,
		CurrentTemperature: 15,
		MinimumTemperature: 8,
		CurrentCapacity:    14,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseID:        5,
		ProductTypeID:      12,
	})
	//assert
	assert.Nil(t, err)
	assert.True(t, repoMock.SaveFlag)
	assert.Equal(t, expSection, result)
}

// test to create section when data initial has information
func TestSectionServiceCreateOkwithData(t *testing.T) {
	//arrange
	expSection := domain.Section{
		ID:                 4,
		SectionNumber:      19,
		CurrentTemperature: 24,
		MinimumTemperature: 15,
		CurrentCapacity:    372,
		MinimumCapacity:    145,
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	}
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      19,
		CurrentTemperature: 24,
		MinimumTemperature: 15,
		CurrentCapacity:    372,
		MinimumCapacity:    145,
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	})
	//assert
	assert.Nil(t, err)
	assert.True(t, repoMock.SaveFlag)
	assert.Equal(t, expSection, result)
}

// test to create section with section number exist in initial data
func TestSectionServiceCreateConflict(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      14,
		CurrentTemperature: 24,
		MinimumTemperature: 15,
		CurrentCapacity:    372,
		MinimumCapacity:    145,
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	})
	//assert
	assert.ErrorIs(t, err, ErrExists)
	assert.False(t, repoMock.SaveFlag)
	assert.Empty(t, result)
}

// test to create section but the current temperature is lower than the minimum.
func TestSectionServiceCreateInvalidCurrentTemperature(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	expErr := ErrInvalidValue{Msg: "the current temperature can't be less than the minimum"}
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      16,
		CurrentTemperature: 8, // sent current temperature less than minimum
		MinimumTemperature: 15,
		CurrentCapacity:    372,
		MinimumCapacity:    145,
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	})
	//assert
	assert.ErrorIs(t, err, expErr)
	assert.False(t, repoMock.SaveFlag)
	assert.Empty(t, result)
}

// test to create section but the minimum capacity is lower than 1.
func TestSectionServiceCreateInvalidMinimumCapacity(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	expErr := ErrInvalidValue{Msg: "the minimum capacity must be greater than zero"}
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      16,
		CurrentTemperature: 8,
		MinimumTemperature: 6,
		CurrentCapacity:    372,
		MinimumCapacity:    0, // sent 0 minimum capacity
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	})
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, expErr.Error(), err.Error())
	assert.False(t, repoMock.SaveFlag)
	assert.Empty(t, result)
}

// test to create section but the current capacity is lower than the minimum.
func TestSectionServiceCreateInvalidCurrentCapacityLessThanMin(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	expErr := ErrInvalidValue{Msg: "the current capacity can't be less than the minimum"}
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      16,
		CurrentTemperature: 22,
		MinimumTemperature: 15,
		CurrentCapacity:    100, // sent current capacity less than minimum
		MinimumCapacity:    145,
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	})
	//assert
	assert.ErrorIs(t, err, expErr)
	assert.False(t, repoMock.SaveFlag)
	assert.Empty(t, result)
}

// test to create section but the current capacity is more than the maximum.
func TestSectionServiceCreateInvalidCurrentCapacityMoreThanMax(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	expErr := ErrInvalidValue{Msg: "the current capacity can't be more than the maximum"}
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:      16,
		CurrentTemperature: 22,
		MinimumTemperature: 15,
		CurrentCapacity:    551, // sent current capacity more than maximum
		MinimumCapacity:    145,
		MaximumCapacity:    550,
		WarehouseID:        6,
		ProductTypeID:      34,
	})
	//assert
	assert.ErrorIs(t, err, expErr)
	assert.False(t, repoMock.SaveFlag)
	assert.Empty(t, result)
}

// test to create section when has error in repository
func TestSectionServiceCreateRepoError(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{ErrorRepo: ErrRepo}
	service := NewService(&repoMock)
	//act
	result, err := service.Create(nil, request.SectionPost{
		SectionNumber:   16,
		MinimumCapacity: 1,
		MaximumCapacity: 1,
		CurrentCapacity: 1,
	})
	//assert
	assert.ErrorIs(t, err, ErrRepo)
	assert.True(t, repoMock.SaveFlag)
	assert.Empty(t, result)
}

// READ ALL
// test to ReadAll section
func TestSectionServiceReadAll(t *testing.T) {
	// arrange
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	// act
	result, err := service.GetAll(nil)
	// assert
	assert.Nil(t, err, ErrRepo)
	assert.True(t, repoMock.GetAllFlag)
	assert.Equal(t, initData, result)
}

// READ BY ID
// test to read, get section by id existent
func TestSectionServiceFindByIDExistent(t *testing.T) {
	// arrange
	expSection := domain.Section{
		ID:                 2,
		SectionNumber:      14,
		CurrentTemperature: 12,
		MinimumTemperature: 8,
		CurrentCapacity:    25,
		MinimumCapacity:    1,
		MaximumCapacity:    250,
		WarehouseID:        7,
		ProductTypeID:      16,
	}
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	// act
	result, err := service.Get(nil, expSection.ID)
	// assert
	assert.Nil(t, err, ErrRepo)
	assert.True(t, repoMock.GetFlag)
	assert.Equal(t, expSection, result)
}

// test to read, get section by id non existent
func TestSectionServiceFindByIDNonExistent(t *testing.T) {
	// arrange
	idNonExistent := 4
	repoMock := section.RepositoryMock{DataDummy: initData}
	service := NewService(&repoMock)
	// act
	result, err := service.Get(nil, idNonExistent)
	// assert
	assert.ErrorIs(t, err, ErrNotFound)
	assert.True(t, repoMock.GetFlag)
	assert.Empty(t, result)
}

// test to read, get section by id with error with the repository
func TestSectionServiceFindByIDErrRepo(t *testing.T) {
	// arrange
	repoMock := section.RepositoryMock{ErrorFind: ErrRepo}
	service := NewService(&repoMock)
	// act
	result, err := service.Get(nil, 0)
	// assert
	assert.ErrorIs(t, err, ErrRepo)
	assert.True(t, repoMock.GetFlag)
	assert.Empty(t, result)
}

// UPDATE
// test to update section that exist, set cuurent temperature and capacity
func TestSectionServiceUpdateExistentSetSomeField(t *testing.T) {
	//arrange
	expSection := domain.Section{
		ID:                 2,
		SectionNumber:      14,
		CurrentTemperature: 25,
		MinimumTemperature: 8,
		CurrentCapacity:    132,
		MinimumCapacity:    1,
		MaximumCapacity:    250,
		WarehouseID:        7,
		ProductTypeID:      16,
	}
	repoMock := section.RepositoryMock{DataDummy: initialDataCopy()}
	service := NewService(&repoMock)
	//act
	result, err := service.Update(nil, expSection.ID, request.SectionPatch{
		CurrentTemperature: &expSection.CurrentTemperature,
		CurrentCapacity:    &expSection.CurrentCapacity,
	})
	//assert
	assert.Nil(t, err)
	assert.True(t, repoMock.UpdateFlag)
	assert.Equal(t, expSection, result)
}

// test to update section that exist, set cuurent temperature and capacity
func TestSectionServiceUpdateExistentSetAllField(t *testing.T) {
	//arrange
	expSection := domain.Section{
		ID:                 2,
		SectionNumber:      21,
		CurrentTemperature: 23,
		MinimumTemperature: 5,
		CurrentCapacity:    246,
		MinimumCapacity:    14,
		MaximumCapacity:    520,
		WarehouseID:        8,
		ProductTypeID:      3,
	}
	repoMock := section.RepositoryMock{DataDummy: initialDataCopy()}
	service := NewService(&repoMock)
	//act
	result, err := service.Update(nil, expSection.ID, request.SectionPatch{
		SectionNumber:      &expSection.SectionNumber,
		CurrentTemperature: &expSection.CurrentTemperature,
		MinimumTemperature: &expSection.MinimumTemperature,
		CurrentCapacity:    &expSection.CurrentCapacity,
		MinimumCapacity:    &expSection.MinimumCapacity,
		MaximumCapacity:    &expSection.MaximumCapacity,
		WarehouseID:        &expSection.WarehouseID,
		ProductTypeID:      &expSection.ProductTypeID,
	})
	//assert
	assert.Nil(t, err)
	assert.True(t, repoMock.UpdateFlag)
	assert.Equal(t, expSection, result)
}

// test to update section that exist, but section number exist in data
func TestSectionServiceUpdateExistentAndInvalidSectionNumber(t *testing.T) {
	//arrange
	sectionInvalid := domain.Section{
		ID:            2,
		SectionNumber: 13,
	}
	repoMock := section.RepositoryMock{DataDummy: initialDataCopy()}
	service := NewService(&repoMock)
	//act
	result, err := service.Update(nil, sectionInvalid.ID, request.SectionPatch{
		SectionNumber: &sectionInvalid.SectionNumber,
	})
	//assert
	assert.ErrorIs(t, err, ErrExists)
	assert.True(t, repoMock.ExistsFlag)
	assert.False(t, repoMock.UpdateFlag)
	assert.Empty(t, result)
}

// test to update section that exist, but section number to change is 0
func TestSectionServiceUpdateExistentAndZeroSectionNumber(t *testing.T) {
	//arrange
	sectionInvalid := domain.Section{
		ID:            2,
		SectionNumber: 0, // sent section number as 0
	}
	repoMock := section.RepositoryMock{DataDummy: initialDataCopy()}
	service := NewService(&repoMock)
	expErr := ErrInvalidValue{Msg: "the section number could not be 0"}
	//act
	result, err := service.Update(nil, sectionInvalid.ID, request.SectionPatch{
		SectionNumber: &sectionInvalid.SectionNumber,
	})
	//assert
	assert.ErrorIs(t, err, expErr)
	assert.True(t, repoMock.ExistsFlag)
	assert.False(t, repoMock.UpdateFlag)
	assert.Empty(t, result)
}

// test to update section that non exist
func TestSectionServiceUpdateNonExistent(t *testing.T) {
	//arrange
	idNonExistent := 4
	repoMock := section.RepositoryMock{}
	service := NewService(&repoMock)
	//act
	result, err := service.Update(nil, idNonExistent, request.SectionPatch{})
	//assert
	assert.ErrorIs(t, err, ErrNotFound)
	assert.False(t, repoMock.UpdateFlag)
	assert.Empty(t, result)
}

// test to update section with error with the repository
func TestSectionServiceUpdateErrRepo(t *testing.T) {
	//arrange
	repoMock := section.RepositoryMock{ErrorRepo: ErrRepo, DataDummy: initData}
	service := NewService(&repoMock)
	//act
	result, err := service.Update(nil, 1, request.SectionPatch{})
	//assert
	assert.ErrorIs(t, err, ErrRepo)
	assert.True(t, repoMock.UpdateFlag)
	assert.Empty(t, result)
}

// DELETE
// test to delete section that exist
func TestSectionServiceDeleteOk(t *testing.T) {
	//arrange
	sectionToDelete := 2
	repoMock := section.RepositoryMock{DataDummy: initialDataCopy()}
	service := NewService(&repoMock)
	//act
	err := service.Delete(nil, sectionToDelete)
	//assert
	assert.Nil(t, err)
	assert.True(t, repoMock.DeleteFlag)
	assert.Equal(t, append(initData[:1], initData[2:]...), repoMock.DataDummy)
}

// test to delete section that exist
func TestSectionServiceDeleteNonExistent(t *testing.T) {
	//arrange
	sectionToDelete := 4
	repoMock := section.RepositoryMock{DataDummy: initialDataCopy()}
	service := NewService(&repoMock)
	//act
	err := service.Delete(nil, sectionToDelete)
	//assert
	assert.ErrorIs(t, err, ErrNotFound)
	assert.True(t, repoMock.DeleteFlag)
	assert.Equal(t, initData, repoMock.DataDummy)
}
