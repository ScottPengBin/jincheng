package base

import "github.com/gin-gonic/gin"

func RegisterGin() *gin.Engine {
	g := gin.Default()
	return g
}
