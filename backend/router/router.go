package router

import (
	"project_management/controller"

	"github.com/gin-gonic/gin"
)

func routers(r *gin.Engine) {
	r.GET("/test", controller.Test)
	task := r.Group("/task")
	task.GET("/tasks", controller.TaskIndex)
	task.POST("/store", controller.TaskStore)
	task.PATCH("/update", controller.TaskUpdate)
	task.DELETE("/delete", controller.TaskDelete)
	box := r.Group("/box")
	box.GET("/tasks", controller.BoxIndex)
	box.POST("/store", controller.BoxStore)
	box.PATCH("/update", controller.BoxUpdate)
	box.DELETE("/delete", controller.BoxDelete)

}
