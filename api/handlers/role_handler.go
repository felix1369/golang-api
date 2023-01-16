package handlers

import (
	"net/http"
	"strconv"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/felix1369/golang-api/model"
	"github.com/felix1369/golang-api/model/entities"
	"github.com/felix1369/golang-api/model/interfaces"
	"github.com/labstack/echo"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// RoleHandler  represent the httphandler for role
type RoleHandler struct {
	RoleUsecase interfaces.RoleUseCase
}

// NewRoleHandler will initialize the role/ resources endpoint
func NewRoleHandler(e *echo.Echo, us interfaces.RoleUseCase) {
	handler := &RoleHandler{
		RoleUsecase: us,
	}
	e.GET("/role", handler.FetchRole)
	e.POST("/role", handler.Store)
	e.GET("/role/:id", handler.GetByID)
	e.DELETE("/role/:id", handler.Delete)
}

func (a *RoleHandler) FetchRole(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := a.RoleUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func (a *RoleHandler) Store(c echo.Context) (err error) {
	var role entities.Role
	err = c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&role); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.RoleUsecase.Store(ctx, &role)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, role)
}

func (a *RoleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	id := uint(idP)
	ctx := c.Request().Context()

	art, err := a.RoleUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func (a *RoleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	id := uint(idP)
	ctx := c.Request().Context()

	err = a.RoleUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(m *entities.Role) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
