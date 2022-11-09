package section

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	r "github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
)

// Errors
var (
	ErrNotFound      = errors.New("section not found")
	ErrExists        = errors.New("section number exists")
	ErrNotValidValue = errors.New("Not valid value")
)

type ErrInvalidValue struct {
	Msg string
}

func (err ErrInvalidValue) Error() string {
	return fmt.Sprintf("not valid value: %s", err.Msg)
}

type Service interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Create(ctx context.Context, request r.SectionPost) (domain.Section, error)
	Update(ctx context.Context, id int, request r.SectionPatch) (domain.Section, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (se *service) GetAll(ctx context.Context) ([]domain.Section, error) {
	return se.repository.GetAll(ctx)
}

func (se *service) Get(ctx context.Context, id int) (domain.Section, error) {
	section, err := se.repository.Get(ctx, id)
	if err != nil {
		return domain.Section{}, switchErrorService(err)
	}
	return section, nil
}

func (se *service) Create(ctx context.Context, request r.SectionPost) (domain.Section, error) {
	switch {
	case se.repository.Exists(ctx, request.SectionNumber):
		return domain.Section{}, ErrExists
	case request.CurrentTemperature < request.MinimumTemperature:
		return domain.Section{}, ErrInvalidValue{Msg: "the current temperature can't be less than the minimum"}
	case request.MinimumCapacity < 1:
		return domain.Section{}, ErrInvalidValue{Msg: "the minimum capacity must be greater than zero"}
	case request.CurrentCapacity < request.MinimumCapacity:
		return domain.Section{}, ErrInvalidValue{Msg: "the current capacity can't be less than the minimum"}
	case request.CurrentCapacity > request.MaximumCapacity:
		return domain.Section{}, ErrInvalidValue{Msg: "the current capacity can't be more than the maximum"}
	}
	section := domain.Section{
		SectionNumber:      request.SectionNumber,
		CurrentTemperature: request.CurrentTemperature,
		MinimumTemperature: request.MinimumTemperature,
		CurrentCapacity:    request.CurrentCapacity,
		MinimumCapacity:    request.MinimumCapacity,
		MaximumCapacity:    request.MaximumCapacity,
		WarehouseID:        request.WarehouseID,
		ProductTypeID:      request.ProductTypeID,
	}
	id, err := se.repository.Save(ctx, section)
	if err != nil {
		return domain.Section{}, err
	}
	section.ID = id
	return section, nil
}

func (se *service) Update(ctx context.Context, id int, request r.SectionPatch) (domain.Section, error) {
	section, err := se.repository.Get(ctx, id)
	if err != nil {
		return domain.Section{}, switchErrorService(err)
	}
	if request.SectionNumber != nil {
		if se.repository.Exists(ctx, *request.SectionNumber) {
			return domain.Section{}, ErrExists
		}
		if *request.SectionNumber == 0 {
			return domain.Section{}, ErrInvalidValue{Msg: "the section number could not be 0"}
		}
		section.SectionNumber = *request.SectionNumber
	}
	if request.MinimumTemperature != nil {
		section.MinimumTemperature = *request.MinimumTemperature
	}
	if request.CurrentTemperature != nil {
		section.CurrentTemperature = *request.CurrentTemperature
	}
	if request.MinimumCapacity != nil {
		section.MinimumCapacity = *request.MinimumCapacity
	}
	if request.MaximumCapacity != nil {
		section.MaximumCapacity = *request.MaximumCapacity
	}
	if request.CurrentCapacity != nil {
		section.CurrentCapacity = *request.CurrentCapacity
	}
	if request.WarehouseID != nil {
		section.WarehouseID = *request.WarehouseID
	}
	if request.ProductTypeID != nil {
		section.ProductTypeID = *request.ProductTypeID
	}
	err = se.repository.Update(ctx, section)
	if err != nil {
		return domain.Section{}, err
	}
	return section, nil
}

func (se *service) Delete(ctx context.Context, id int) error {
	err := se.repository.Delete(ctx, id)
	if err != nil {
		return switchErrorService(err)
	}
	return nil
}

func switchErrorService(err error) error {
	if err == sql.ErrNoRows {
		return ErrNotFound
	}
	return err
}
