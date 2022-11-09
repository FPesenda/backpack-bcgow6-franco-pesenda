package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w6-4/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w6-4/pkg/web/request"
	"github.com/gin-gonic/gin"
)

type Employee struct {
	employeeService employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

// ListEmployees godoc
// @Summary     Get employees by ID
// @Tags        Employees
// @Description get employees
// @Produce     json
// @Param       token header   string true "token"
// @Param       id    path     string true "employee id"
// @Success     200   {object} web.response
// @Failure     404   {object} web.errorResponse
// @Router      /employee/{id} [get]
func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: invalid id - %d", id)
			return
		}
		emp, err := e.employeeService.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusOK, emp)
	}
}

// ListEmployees godoc
// @Summary     List employees
// @Tags        Employees
// @Description get employees
// @Produce     json
// @Param       token header   string true "token"
// @Success     200   {object} web.response
// @Failure     400   {object} web.errorResponse
// @Router      /employee [get]
func (e *Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		emp, err := e.employeeService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}
		if len(emp) == 0 {
			web.Error(c, http.StatusBadRequest, "error: %s", "there are no registered employees")
			return
		}
		web.Success(c, http.StatusOK, emp)
	}
}

// CreateEmployee godoc
// @Summary     Create a new employee
// @Tags        Employees
// @Description create a new employee with all attributes
// @accept      json
// @Produce     json
// @Param       token    header   string                 true "token"
// @Param       employee body     request.EmployeeCreate true "employee to create"
// @Success     201      {object} web.response
// @Failure     422      {object} web.errorResponse
// @Router      /employee [post]
func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.EmployeeCreate
		if err := c.ShouldBindJSON(&request); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error: %s", err.Error())
			return
		}
		id, err := e.employeeService.Save(c, request)
		if err != nil {
			switchErrorEmployee(c, err)
			return
		}
		web.Success(c, http.StatusCreated, id)
	}
}

// UpdateEmployee godoc
// @Summary     Update a employee
// @Tags        Employees
// @Description update a employee selected by id
// @accept      json
// @Produce     json
// @Param       token    header   string                true "token"
// @Param       id       path     string                true "employee id"
// @Param       employee body     request.EmployeePatch true "employee to update"
// @Success     200      {object} web.response
// @Failure     400      {object} web.errorResponse
// @Failure     404      {object} web.errorResponse
// @Router      /employee/{id} [patch]
func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: invalid id - %d", id)
			return
		}
		var empPatch request.EmployeePatch
		if err := c.ShouldBindJSON(&empPatch); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}
		empDB, err := e.employeeService.Update(c, id, empPatch)
		if err != nil {
			switchErrorEmployee(c, err)
			return
		}
		web.Success(c, http.StatusOK, empDB)
	}
}

// DeleteSection godoc
// @Summary     Delete a employee
// @Tags        Employees
// @Description Delete a employee selected by id
// @Param       token header   string true "token"
// @Param       id    path     string true "employee id"
// @Success     204   {object} web.response
// @Failure     404   {object} web.errorResponse
// @Router      /employee/{id} [delete]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusNotFound, "error: invalid id - %d", id)
			return
		}
		if err := e.employeeService.Delete(c, id); err != nil {
			web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}

func switchErrorEmployee(c *gin.Context, err error) {
	switch err {
	case employee.ErrNotFound:
		web.Error(c, http.StatusNotFound, "error %s", err.Error())
		return
	case employee.ErrExist:
		web.Error(c, http.StatusConflict, "error %s", err.Error())
		return
	default:
		web.Error(c, http.StatusInternalServerError, "error %s", err.Error())
		return
	}
}
