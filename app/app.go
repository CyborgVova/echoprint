package app

import (
	"fmt"
	"net/http"

	"github.com/cyborgvova/echoprint/config"
	"github.com/gin-gonic/gin"
)

type App struct {
	e   *gin.Engine
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	e := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": cfg.Text})
	})

	e.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ready"})
	})

	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "health"})
	})

	return &App{
		e:   e,
		cfg: cfg,
	}
}

func (a *App) Handler() http.Handler {
	return a.e.Handler()
}

func (a *App) Start() error {
	return a.e.Run(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *App) Stop() error {
	return nil
}
