package users

import (
	"context"
	"net/http"

	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(ctx context.Context, req *request.CreateUser) (*response.User, error)
	ValidateCredentials(ctx context.Context, req *request.ValidateUserCredentials) (*response.User, error)
}

type Handler struct {
	userService UserService
}

func NewHandler(userService UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req request.CreateUser

	if err := c.Bind(&req); err != nil {
		return errors.New(
			http.StatusBadRequest,
			errors.StatusBadRequestCode,
			[]string{"Invalid request body"},
		)
	}

	result, err := h.userService.CreateUser(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *Handler) ValidateCredentials(c echo.Context) error {
	var req request.ValidateUserCredentials

	if err := c.Bind(&req); err != nil {
		return errors.New(
			http.StatusBadRequest,
			errors.StatusBadRequestCode,
			[]string{"Invalid request body"},
		)
	}

	result, err := h.userService.ValidateCredentials(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
