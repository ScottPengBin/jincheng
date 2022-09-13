package middle_ware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LogMiddleWare(logger *logrus.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		st := time.Now()

		c.Next()

		latencyTime := fmt.Sprintf("%fs", time.Now().Sub(st).Seconds())

		logger.WithFields(logrus.Fields{
			"latency_time": latencyTime,
			"method":       c.Request.Method,
			"req_url":      c.Request.RequestURI,
			"client_id":    c.ClientIP(),
			"status_code":  c.Writer.Status(),
		}).Info()
	}

}
