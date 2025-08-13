package helpers

import (
	"github.com/juanMaAV92/go-utils/log"
	"github.com/juanMaAV92/zenith-financial/backend/cmd"
	"github.com/juanMaAV92/zenith-financial/backend/platform/config"
)

func NewTestServer() *cmd.Instance {
	cfg, _ := config.Load("local")
	logger := log.New(config.MicroserviceName, log.WithLevel(log.DebugLevel))
	testServer, _ := cmd.NewServer(cfg, logger)
	return testServer
}
