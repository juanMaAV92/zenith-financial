package config

import (
	"github.com/juanMaAV92/go-utils/database"
	platform "github.com/juanMaAV92/go-utils/platform/config"
)

type Config struct {
	*platform.BasicConfig
	Telemetry platform.TelemetryConfig
	Database  database.DBConfig
}
