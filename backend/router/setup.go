package router

import (
	"project_management/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	gin.SetMode(gin.DebugMode)
	routers(r)
	r.Run(":8090")
}
