package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jincheng/internal/config"
	"jincheng/internal/middle_ware"
	"net/http"
)

type Server struct {
	host       string
	port       string
	logger     logrus.Logger
	router     gin.Engine
	httpServer http.Server
}

func newHttpServer(config config.Config, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":" + config.App.Port,
		Handler: router,
	}
}

func NewRouter(config config.Config, logger *logrus.Logger) *gin.Engine {
	gin.SetMode(config.App.Mode)

	r := gin.New()

	r.Use(middle_ware.LogMiddleWare(logger))

	return r
}
