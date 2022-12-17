package router

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	routers(r)
	r.Run(":8090")
}
