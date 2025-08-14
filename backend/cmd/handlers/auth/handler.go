package auth

import (
	"context"
	"net/http"

	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/go-utils/headers"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/request"
	"github.com/juanMaAV92/zenith-financial/backend/internal/domain/response"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Login(ctx context.Context, req *request.UserLogin) (*response.UserLogin, error)
	Logout(ctx context.Context, authHeader string) error
	RefreshToken(ctx context.Context, refreshToekn string) (*response.TokensResponse, error)
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

func (h *Handler) Logout(c echo.Context) error {
	var authHeader string
	if authHeader = c.Request().Header.Get(headers.Authorization); authHeader == "" {
		return errors.New(
			http.StatusUnauthorized,
			errors.StatusUnauthorizedCode,
			[]string{"Authorization header is required"},
		)
	}

	err := h.authService.Logout(c.Request().Context(), authHeader)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var req request.RefreshToken

	if err := c.Bind(&req); err != nil {
		return errors.New(
			http.StatusBadRequest,
			errors.StatusBadRequestCode,
			[]string{"Invalid request body"},
		)
	}

	result, err := h.authService.RefreshToken(c.Request().Context(), req.RefreshToken)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
