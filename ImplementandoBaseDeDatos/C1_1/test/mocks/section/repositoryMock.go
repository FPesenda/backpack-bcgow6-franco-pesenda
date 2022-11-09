package section

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
)

type RepositoryMock struct {
	DataDummy  []domain.Section
	ErrorRepo  error
	ErrorFind  error
	GetAllFlag bool
	GetFlag    bool
	ExistsFlag bool
	SaveFlag   bool
	UpdateFlag bool
	DeleteFlag bool
}

func (r *RepositoryMock) GetAll(ctx context.Context) ([]domain.Section, error) {
	r.GetAllFlag = true
	if r.ErrorRepo != nil {
		return nil, r.ErrorRepo
	}
	return r.DataDummy, nil
}
func (r *RepositoryMock) Get(ctx context.Context, id int) (domain.Section, error) {
	r.GetFlag = true
	if r.ErrorFind != nil {
		return domain.Section{}, r.ErrorFind
	}
	for _, section := range r.DataDummy {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, sql.ErrNoRows
}
func (r *RepositoryMock) Exists(ctx context.Context, cid int) bool {
	r.ExistsFlag = true
	for _, section := range r.DataDummy {
		if section.SectionNumber == cid {
			return true
		}
	}
	return false
}
func (r *RepositoryMock) Save(ctx context.Context, s domain.Section) (int, error) {
	r.SaveFlag = true
	if r.ErrorRepo != nil {
		return 0, r.ErrorRepo
	}
	s.ID = r.nextID()
	r.DataDummy = append(r.DataDummy, s)
	return s.ID, nil
}
func (r *RepositoryMock) Update(ctx context.Context, s domain.Section) error {
	r.UpdateFlag = true
	if r.ErrorRepo != nil {
		return r.ErrorRepo
	}
	for i := range r.DataDummy {
		if r.DataDummy[i].ID == s.ID {
			r.DataDummy[i] = s
			return nil
		}
	}
	return sql.ErrNoRows
}
func (r *RepositoryMock) Delete(ctx context.Context, id int) error {
	r.DeleteFlag = true
	if r.ErrorRepo != nil {
		return r.ErrorRepo
	}
	for i := range r.DataDummy {
		if r.DataDummy[i].ID == id {
			r.DataDummy = append(r.DataDummy[:i], r.DataDummy[i+1:]...)
			return nil
		}
	}
	return sql.ErrNoRows
}
func (r *RepositoryMock) nextID() int {
	if r.DataDummy == nil || len(r.DataDummy) == 0 {
		return 1
	}
	return r.DataDummy[len(r.DataDummy)-1].ID + 1
}
