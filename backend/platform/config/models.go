package config

import (
	"github.com/juanMaAV92/go-utils/cache"
	"github.com/juanMaAV92/go-utils/database"
	"github.com/juanMaAV92/go-utils/jwt"
	platform "github.com/juanMaAV92/go-utils/platform/config"
)

type Config struct {
	*platform.BasicConfig
	Telemetry *platform.TelemetryConfig
	Database  *database.DBConfig
	Jwt       *jwt.JwtConfig
	Cache     *cache.CacheConfig
}
