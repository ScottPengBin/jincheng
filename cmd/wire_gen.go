// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	admin2 "jincheng/app/controller/admin"
	"jincheng/app/controller/member"
	"jincheng/config"
	"jincheng/internal/core/base/http"
	"jincheng/internal/core/db"
	"jincheng/internal/core/log"
	"jincheng/internal/router"
	"jincheng/internal/service/admin"
	"jincheng/internal/service/member"
	http2 "net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Injectors from wire.go:

func InitApp() *App {
	configConfig := config.GetConfig()
	logger := log.NewLog()
	dataBase := db.NewDataBase(configConfig, logger)
	service := memberSer.NewService(dataBase)
	controller := member.NewController(service)
	adminService := admin.NewService(dataBase)
	adminController := admin2.NewController(adminService)
	optionsController := &router.OptionsController{
		Member: controller,
		Admin:  adminController,
	}
	v := router.Router(optionsController)
	engine := http.NewRouter(configConfig, logger, v)
	myServer := http.NewHttpServer(configConfig, engine)
	app := &App{
		DataBase:   dataBase,
		Config:     configConfig,
		Logger:     logger,
		Router:     engine,
		HttpServer: myServer,
	}
	return app
}

// wire.go:

type App struct {
	DataBase   *db.DataBase
	Config     *config.Config
	Logger     *logrus.Logger
	Router     *gin.Engine
	HttpServer *http.MyServer
	mu         sync.RWMutex
	done       []chan os.Signal
}

func (app *App) Start() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	app.mu.Lock()
	go func() {
		err := app.HttpServer.Start()
		if err != nil && err != http2.ErrServerClosed {
			panic(err)
		}
	}()
	app.done = append(app.done, c)
	app.mu.Unlock()
	return c
}

func (app *App) Stop() error {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		app.HttpServer.Stop()
		app.Logger.Info("关闭http")
		defer wg.Done()
	}()

	go func() {
		mdbStats := app.DataBase.MasterDb.Stats()
		sdbStats := app.DataBase.SalveDb.Stats()
		i := 0
		for {
			if mdbStats.Idle == 0 {
				app.DataBase.MasterDb.Close()
				app.Logger.Info("关闭MasterDb")
				i++
			}
			if sdbStats.Idle == 0 {
				app.DataBase.SalveDb.Close()
				app.Logger.Info("关闭SalveDb")
				i++
			}
			if i >= 2 {
				break
			}
			time.Sleep(time.Second)
		}
		defer wg.Done()
	}()

	wg.Wait()

	return nil
}
