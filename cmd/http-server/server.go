package main

import (
	"os"

	// "github.com/illacloud/builder-backend/pkg/cors"
	// "github.com/illacloud/builder-backend/pkg/recovery"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sjxiang/hole/api/router"
)

type Config struct {
	HOLE_SERVER_HOST string `env:"HOLE_SERVER_HOST" envDefault:"0.0.0.0"`
	HOLE_SERVER_PORT string `env:"HOLE_SERVER_PORT" envDefault:"8001"`
	HOLE_SERVER_MODE string `env:"HOLE_SERVER_MODE" envDefault:"debug"`
}



func GetAppConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type Server struct {
	engine     *gin.Engine
	restRouter *router.RESTRouter
	logger     *zap.SugaredLogger
	cfg        *Config
}

func NewServer(cfg *Config, engine *gin.Engine, restRouter *router.RESTRouter, logger *zap.SugaredLogger) *Server {
	return &Server{
		engine:     engine,
		cfg:        cfg,
		restRouter: restRouter,
		logger:     logger,
	}
}

func (server *Server) Start() {
	server.logger.Infow("Starting server")

	gin.SetMode(server.cfg.HOLE_SERVER_MODE)

	// 全局中间件
	// corsHandleRecovery := recovery.CorsHandleRecovery()
	// server.engine.Use(gin.CustomRecovery(corsHandleRecovery))
	// server.engine.Use(cors.Cors())

	server.restRouter.InitRouter(server.engine.Group("/api"))

	err := server.engine.Run(server.cfg.HOLE_SERVER_HOST + ":" + server.cfg.HOLE_SERVER_PORT)
	if err != nil {
		server.logger.Errorw("Error in startup", "err", err)
		os.Exit(2)
	}
}
