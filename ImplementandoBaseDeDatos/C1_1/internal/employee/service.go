package employee

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
)

// Errors
var (
	ErrCardNumberID = errors.New("card number id cannot be updated")
	ErrNotFound     = errors.New("employee not found")
	ErrExist        = errors.New("existing resource")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, request request.EmployeeCreate) (domain.Employee, error)
	Update(ctx context.Context, id int, request request.EmployeePatch) (domain.Employee, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}
func (s *service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return s.repository.GetAll(ctx)
}
func (s *service) Get(ctx context.Context, id int) (domain.Employee, error) {
	return s.repository.Get(ctx, id)
}
func (s *service) Exists(ctx context.Context, cardNumberID string) bool {
	return s.repository.Exists(ctx, cardNumberID)
}
func (s *service) Save(ctx context.Context, request request.EmployeeCreate) (domain.Employee, error) {
	if s.repository.Exists(ctx, request.CardNumberID) {
		return domain.Employee{}, ErrExist
	}
	emp := domain.Employee{
		CardNumberID: request.CardNumberID,
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		WarehouseID:  request.WarehouseID,
	}
	id, err := s.repository.Save(ctx, emp)
	if err != nil {
		return domain.Employee{}, ErrExist
	}
	emp.ID = id
	return emp, nil
}
func (s *service) Update(ctx context.Context, id int, request request.EmployeePatch) (domain.Employee, error) {
	empDB, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Employee{}, ErrNotFound
	}
	// FirstName
	if request.FirstName != nil {
		empDB.FirstName = *request.FirstName
	}
	// LastName
	if request.LastName != nil {
		empDB.LastName = *request.LastName
	}
	// WarehouseID
	if request.WarehouseID != nil {
		empDB.WarehouseID = *request.WarehouseID
	}
	err = s.repository.Update(ctx, empDB)
	if err != nil {
		return domain.Employee{}, err
	}
	return empDB, nil
}
func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
