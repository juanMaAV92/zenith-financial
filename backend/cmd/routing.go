package cmd

import (
	utilsMiddleware "github.com/juanMaAV92/go-utils/middleware"
	"github.com/juanMaAV92/zenith-financial/backend/cmd/handlers/health"
	"github.com/juanMaAV92/zenith-financial/backend/cmd/handlers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiV1Group      = "/v1"
	healthCheckPath = "/health-check"
	userPath        = "/users"
	loginPath       = "/users/login"
	registerPath    = "/users/register"
)

type HealthHandler interface {
	Check(ctx echo.Context) error
}

type UserHandler interface {
	CreateUser(ctx echo.Context) error
	ValidateCredentials(ctx echo.Context) error
}

type handlers struct {
	health HealthHandler
	user   UserHandler
}

func configRoutes(inst *Instance, services *services) {
	configMiddleware(inst)

	handlers := initializeHandlers(services)

	baseGroup := inst.Server.Group(inst.config.ServerName)
	baseGroup.GET(healthCheckPath, handlers.health.Check)

	v1 := baseGroup.Group(apiV1Group)
	configureV1Routes(v1, handlers)
}

func initializeHandlers(services *services) *handlers {
	healthHandler := health.NewHandler(services.healthService)
	UserHandler := users.NewHandler(services.userService)

	return &handlers{
		health: healthHandler,
		user:   UserHandler,
	}
}

func configureV1Routes(v1 *echo.Group, h *handlers) {
	v1.POST(registerPath, h.user.CreateUser)
	v1.POST(loginPath, h.user.ValidateCredentials)
}

func configMiddleware(inst *Instance) {
	inst.Server.Use(middleware.Recover())

	inst.Server.Use(utilsMiddleware.Tracing(inst.config.ServerName))
	inst.Server.Use(utilsMiddleware.TraceId())
	inst.Server.Use(utilsMiddleware.Logging(inst.Logger))
}
