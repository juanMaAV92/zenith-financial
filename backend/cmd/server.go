package cmd

import (
	"github.com/juanMaAV92/go-utils/database"
	"github.com/juanMaAV92/go-utils/env"
	"github.com/juanMaAV92/go-utils/errors"
	"github.com/juanMaAV92/go-utils/log"
	"github.com/juanMaAV92/go-utils/platform/server"
	healthHandler "github.com/juanMaAV92/zenith-financial/backend/cmd/handlers/health"
	"github.com/juanMaAV92/zenith-financial/backend/internal/services/health"
	"github.com/juanMaAV92/zenith-financial/backend/platform/config"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type Instance struct {
	*server.Server
	log.Logger
	config *config.Config
}

type services struct {
	healthService healthHandler.Service
}

func NewServer(cfg *config.Config, logger log.Logger) (*Instance, error) {
	instance := echo.New()
	instance.HideBanner = true
	instance.HTTPErrorHandler = errors.CustomHTTPErrorHandler
	instance.Debug = cfg.Environment == env.LocalEnvironment
	decimal.MarshalJSONWithoutQuotes = true

	return &Instance{
		Server: &server.Server{
			Echo: instance,
		},
		Logger: logger,
		config: cfg,
	}, nil
}

func (inst Instance) initServices() (*services, error) {
	healthService := health.NewService()

	_, err := database.New(inst.config.Database, inst.Logger)
	if err != nil {
		return nil, err
	}

	return &services{
		healthService: healthService,
	}, nil
}
