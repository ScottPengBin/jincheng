package db

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"jincheng/config"
	"log"
	"time"
)

type DataBase struct {
	Master   gorm.DB
	MasterDb *sql.DB
	Salve    gorm.DB
	SalveDb  *sql.DB
}

var Provider = wire.NewSet(NewDataBase)

func NewDataBase(c config.Config, logger *logrus.Logger) DataBase {

	m, mdb := getConn(c.MySQLConf.Driver, c.MySQLConf.Master.Dsn, c.MySQLConf.Prefix, logger)
	s, sdb := getConn(c.MySQLConf.Driver, c.MySQLConf.Slave.Dsn, c.MySQLConf.Prefix, logger)

	return DataBase{
		Master:   m,
		MasterDb: mdb,
		Salve:    s,
		SalveDb:  sdb,
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

func getConn(driver, dsn, prefix string, logger *logrus.Logger) (gorm.DB, *sql.DB) {
	director := mysql.New(mysql.Config{
		DriverName: driver,
		DSN:        dsn,
	})

	connect, err := gorm.Open(director, &gorm.Config{
		Logger: getDBLogger(logger),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix,
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	connect.Debug()
	db, err := connect.DB()
	if err != nil {
		panic("connect db server failed.")
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxOpenConns(100)
	logger.Info("gorm init")

	return *connect, db

}
