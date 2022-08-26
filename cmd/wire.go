//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"study_wire/config"
	"study_wire/internal/core/log"
	"study_wire/internal/db"
)

type App struct {
	DataBase db.DataBase
	Config   config.Config
	Logger   *logrus.Logger
}

func InitApp() App {
	wire.Build(db.Provider,
		config.Provider,
		log.Provider,
		wire.Struct(new(App), "*"),
	)
	return App{}
}

func (a App) Start() {

}
