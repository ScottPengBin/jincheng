package log

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Provider = wire.NewSet(NewLog)

func NewLog() *logrus.Logger {

	log := logrus.New()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(dir+"/logs/", 0777)
	if err != nil {
		return nil
	}

	file, err := os.OpenFile(dir+"/logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Info("Failed to log to file, using default stderr")
	}
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	output := io.MultiWriter([]io.Writer{
		file,
		os.Stdout,
	}...)
	log.SetOutput(output)
	return log
}
