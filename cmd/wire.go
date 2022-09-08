//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"jincheng/app/controller"
	"jincheng/internal/config"
	"jincheng/internal/core/base/http"
	"jincheng/internal/core/log"
	"jincheng/internal/db"
	http2 "net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	DataBase   db.DataBase
	Config     config.Config
	Logger     *logrus.Logger
	Router     *gin.Engine
	HttpServer *http.MyServer
	mu         sync.RWMutex
	done       []chan os.Signal
}

func InitApp() *App {
	wire.Build(
		db.Provider,
		config.Provider,
		log.Provider,
		http.Provider,
		controller.Provider,
		wire.Struct(new(App), "DataBase", "Config", "Logger", "Router", "HttpServer"),
	)
	return &App{}
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
