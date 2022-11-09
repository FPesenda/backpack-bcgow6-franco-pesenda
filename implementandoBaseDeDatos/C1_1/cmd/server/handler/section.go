package handler

import (
	"errors"
	"net/http"

	sectionP "github.com/extmatperez/meli_bootcamp_go_w6-4/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web"
	r "github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
	"github.com/gin-gonic/gin"
)

type Section struct {
	sectionService sectionP.Service
}

func NewSection(s sectionP.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

// ListSections godoc
// @Summary     List sections
// @Tags        Sections
// @Description get sections
// @Produce     json
// @Param       token header   string true "token"
// @Success     200   {object} web.response
// @failure     500   {object} web.errorResponse
// @Router      /sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sections, err := s.sectionService.GetAll(c)
		if err != nil {
			switchErrorSection(c, err)
			return
		}
		if sections == nil {
			web.Success(c, http.StatusOK, "the data base has no records")
			return
		}
		web.Success(c, http.StatusOK, sections)
	}
}

// SectionByID godoc
// @Summary     Get section by ID
// @Tags        Sections
// @Description get section by id
// @Produce     json
// @Param       token header   string true "token"
// @Param       id    path     string true "Section id"
// @Success     200   {object} web.response
// @failure     404   {object} web.errorResponse
// @failure     500   {object} web.errorResponse
// @Router      /sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		section, err := s.sectionService.Get(c, id)
		if err != nil {
			switchErrorSection(c, err)
			return
		}
		web.Success(c, http.StatusOK, section)
	}
}

// CreateSection godoc
// @Summary     Create a new section
// @Tags        Sections
// @Description create a new section with all attributes
// @accept      json
// @Produce     json
// @Param       token   header   string              true "token"
// @Param       section body     request.SectionPost true "section to create"
// @Success     201     {object} web.response
// @failure     409     {object} web.errorResponse
// @failure     422     {object} web.errorResponse
// @failure     500     {object} web.errorResponse
// @Router      /sections/ [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request r.SectionPost
		if err := c.ShouldBindJSON(&request); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error %s", err.Error())
			return
		}
		section, err := s.sectionService.Create(c, request)
		if err != nil {
			switchErrorSection(c, err)
			return

		}
		web.Success(c, http.StatusCreated, section)
	}
}

// UpdateSection godoc
// @Summary     Update a section
// @Tags        Sections
// @Description update a section selected by id
// @accept      json
// @Produce     json
// @Param       token   header   string               true "token"
// @Param       id      path     string               true "Section id"
// @Param       section body     request.SectionPatch true "section to update"
// @Success     200     {object} web.response
// @failure     404     {object} web.errorResponse
// @failure     409     {object} web.errorResponse
// @failure     422     {object} web.errorResponse
// @failure     500     {object} web.errorResponse
// @Router      /sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request r.SectionPatch
		if err := c.ShouldBindJSON(&request); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error %s", err.Error())
			return
		}
		id := c.GetInt("id_validated")
		section, err := s.sectionService.Update(c, id, request)
		if err != nil {
			switchErrorSection(c, err)
			return
		}
		web.Success(c, http.StatusOK, section)
	}
}

// DeleteSection godoc
// @Summary     Delete a section
// @Tags        Sections
// @Description Delete a section selected by id
// @Param       token header   string true "token"
// @Param       id    path     string true "Section id"
// @Success     204   {object} web.response
// @failure     404   {object} web.errorResponse
// @failure     500   {object} web.errorResponse
// @Router      /sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id_validated")
		err := s.sectionService.Delete(c, id)
		if err != nil {
			switchErrorSection(c, err)
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}

func switchErrorSection(c *gin.Context, err error) {
	switch {
	case errors.Is(err, sectionP.ErrNotFound):
		web.Error(c, http.StatusNotFound, "error %s", err.Error())
		return
	case errors.Is(err, sectionP.ErrExists):
		web.Error(c, http.StatusConflict, "error %s", err.Error())
		return
	case errors.As(err, &sectionP.ErrInvalidValue{}):
		web.Error(c, http.StatusUnprocessableEntity, "error %s", err.Error())
		return
	default:
		web.Error(c, http.StatusInternalServerError, "error %s", err.Error())
		return
	}
}
