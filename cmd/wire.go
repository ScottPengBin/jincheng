//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"jincheng/internal/config"
	"jincheng/internal/core/log"
	"jincheng/internal/db"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	DataBase db.DataBase
	Config   config.Config
	Logger   *logrus.Logger
}

func InitApp() App {
	wire.Build(
		db.Provider,
		config.Provider,
		log.Provider,
		wire.Struct(new(App), "*"),
	)
	return App{}
}

func (app *App) Start() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}
