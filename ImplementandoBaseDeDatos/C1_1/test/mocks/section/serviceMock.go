package section

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/domain"
	r "github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
)

type ServiceMock struct {
	SectionsDummy     []domain.Section
	SectionDummy      domain.Section
	ServiceMethodFlag bool
	ErrService        error
}

func (r *ServiceMock) GetAll(ctx context.Context) ([]domain.Section, error) {
	r.ServiceMethodFlag = true
	if r.ErrService != nil {
		return nil, r.ErrService
	}
	return r.SectionsDummy, nil
}
func (r *ServiceMock) Get(ctx context.Context, id int) (domain.Section, error) {
	r.ServiceMethodFlag = true
	if r.ErrService != nil {
		return domain.Section{}, r.ErrService
	}
	return r.SectionDummy, nil
}
func (r *ServiceMock) Create(ctx context.Context, s r.SectionPost) (domain.Section, error) {
	r.ServiceMethodFlag = true
	if r.ErrService != nil {
		return domain.Section{}, r.ErrService
	}
	return r.SectionDummy, nil
}
func (r *ServiceMock) Update(ctx context.Context, id int, s r.SectionPatch) (domain.Section, error) {
	r.ServiceMethodFlag = true
	if r.ErrService != nil {
		return domain.Section{}, r.ErrService
	}
	return r.SectionDummy, nil
}
func (r *ServiceMock) Delete(ctx context.Context, id int) error {
	r.ServiceMethodFlag = true
	if r.ErrService != nil {
		return r.ErrService
	}
	return nil
}
