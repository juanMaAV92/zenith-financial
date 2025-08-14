package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juanMaAV92/go-utils/cache"
	"github.com/juanMaAV92/go-utils/database"
	"github.com/juanMaAV92/go-utils/env"
	jwtUtils "github.com/juanMaAV92/go-utils/jwt"
	platform "github.com/juanMaAV92/go-utils/platform/config"
	"gorm.io/gorm/logger"
)

const (
	MicroserviceName = "zenith-financial"
	ServiceVersion   = "1.0.0"
)

var localConfig = Config{
	BasicConfig: &platform.BasicConfig{
		Port:         "8080",
		GracefulTime: 5 * time.Minute,
		Environment:  "local",
		ServerName:   MicroserviceName,
	},
	Telemetry: &platform.TelemetryConfig{
		OTLPEndpoint: "localhost:4318",
	},
	Database: &database.DBConfig{
		Host:        "localhost",
		Password:    "postgres",
		User:        "postgres",
		Port:        "5432",
		Name:        fmt.Sprintf("%s-db", MicroserviceName),
		LogLevel:    logger.Info,
		MaxPoolSize: 15,
		MaxLifeTime: 5 * time.Minute,
	},
	Jwt: &jwtUtils.JwtConfig{
		SecretKey:       "secret",
		AccessTokenTTL:  1 * time.Minute,
		RefreshTokenTTL: 24 * time.Hour,
		Issuer:          MicroserviceName,
		SigningMethod:   jwt.SigningMethodHS256,
	},
	Cache: &cache.CacheConfig{
		Host:       "localhost",
		Port:       "6379",
		ServerName: MicroserviceName,
	},
}

func deployConfig() Config {
	return Config{
		BasicConfig: platform.GetBasicServerConfig(MicroserviceName),
		Telemetry: &platform.TelemetryConfig{
			OTLPEndpoint: env.GetEnv(env.OTLP_ENDPOINT),
		},
		Database: database.GetDBConfig(),
		Jwt:      jwtUtils.GetJWTConfig(MicroserviceName, jwt.SigningMethodHS256),
		Cache:    cache.GetCacheConfig(MicroserviceName),
	}
}

func Load(enviroment string) (*Config, error) {
	var config Config
	if enviroment == env.LocalEnvironment {
		config = localConfig
	} else {
		config = deployConfig()
	}

	return &config, nil
}
