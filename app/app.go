package app

import (
	"github.com/cyborgvova/echoprint/config"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	application := &App{
		cfg: cfg,
	}

	return application
}
