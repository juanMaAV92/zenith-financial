package helpers

import (
	"github.com/juanMaAV92/zenith-financial/cmd"
	"github.com/juanMaAV92/zenith-financial/platform/config"
)

func NewTestServer() *cmd.Instance {
	cfg, _ := config.Load("local")
	testServer, _ := cmd.NewServer(cfg, nil)
	return testServer
}
