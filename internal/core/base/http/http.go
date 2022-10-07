package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"jincheng/config"
	"jincheng/internal/core/middle_ware"
	"jincheng/internal/core/valida"
	"jincheng/internal/router"
	"net/http"
	"time"
)

var Provider = wire.NewSet(NewRouter, NewHttpServer, router.Router, router.Provider)

type MyServer struct {
	httpServer *http.Server
}

func NewRouter(config *config.Config, logger *logrus.Logger, controllers func(r *gin.Engine)) *gin.Engine {
	gin.SetMode(config.App.Mode)

	r := gin.New()

	r.Use(middle_ware.LogMiddleWare(logger))
	r.Use(middle_ware.Cors)
	r.Use(middle_ware.Jwt(config))

	controllers(r)

	//加载中文验证器
	valida.LoadZh()

	return r
}

func NewHttpServer(config *config.Config, router *gin.Engine) *MyServer {
	return &MyServer{
		httpServer: &http.Server{
			Addr:    ":" + config.App.Port,
			Handler: router,
		},
	}
}

func (s *MyServer) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *MyServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
