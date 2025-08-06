package auth

import (
	"context"
	"net/http"

	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Login(ctx context.Context, req *request.UserLogin) (*response.UserLogin, error)
}

type Handler struct {
	authService AuthService
}

func NewHandler(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Login(c echo.Context) error {
	var req request.UserLogin

	if err := c.Bind(&req); err != nil {
		return errors.New(
			http.StatusBadRequest,
			errors.StatusBadRequestCode,
			[]string{"Invalid request body"},
		)
	}

	result, err := h.authService.Login(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
