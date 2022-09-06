package db

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jincheng/internal/config"
	"log"
	"time"
)

type DataBase struct {
	Master gorm.DB
	Salve  gorm.DB
}

var Provider = wire.NewSet(NewDataBase)

func NewDataBase(c config.Config, logger *logrus.Logger) DataBase {
	return DataBase{
		Master: getConn(c.MySQLConf.Driver, c.MySQLConf.Master.Dsn, logger),
		Salve:  getConn(c.MySQLConf.Driver, c.MySQLConf.Slave.Dsn, logger),
	}
}

func getDBLogger(l *logrus.Logger) logger.Interface {
	return logger.New(
		log.New(l.Out, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.LogLevel(l.Level),
		},
	)

}

func getConn(driver, dsn string, logger *logrus.Logger) gorm.DB {
	director := mysql.New(mysql.Config{
		DriverName: driver,
		DSN:        dsn,
	})
	connect, _ := gorm.Open(director, &gorm.Config{
		Logger: getDBLogger(logger),
	})
	connect.Debug()
	db, err := connect.DB()
	if err != nil {
		panic("connect db server failed.")
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxOpenConns(100)
	logger.Info("gorm init")
	return *connect

}
