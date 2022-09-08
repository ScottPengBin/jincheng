package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"jincheng/app/pers_amount_total"
	"jincheng/internal/config"
	"jincheng/internal/middle_ware"
	"net/http"
	"time"
)

var Provider = wire.NewSet(NewRouter, NewHttpServer)

type MyServer struct {
	httpServer *http.Server
}

func NewRouter(config config.Config, logger *logrus.Logger,controllers pers_amount_total.InitControllers) *gin.Engine {
	gin.SetMode(config.App.Mode)

	r := gin.New()

	r.Use(middle_ware.LogMiddleWare(logger))

	controllers(r)

	return r
}

func NewHttpServer(config config.Config, router *gin.Engine) *MyServer {
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