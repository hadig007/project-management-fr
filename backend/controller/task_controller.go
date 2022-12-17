package controller

import (
	"log"
	"project_management/database"
	model "project_management/models"

	"github.com/gin-gonic/gin"
)

var task model.TaskModel
var tasks []model.TaskModel

func TaskStore(c *gin.Context) {
	c.ShouldBindJSON(&task)

	result := database.Mysql.Save(&model.TaskModel{
		Name:        task.Name,
		Status:      task.Status,
		Color:       task.Color,
		Tag:         task.Tag,
		Group:       task.Group,
		Description: task.Description,
	})
	if result.Error != nil {
		log.Fatal("Failed save data")
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
	})

}

func TaskIndex(c *gin.Context) {
	data := database.Mysql.Find(&tasks)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}

func TaskUpdate(c *gin.Context) {
	var t model.TaskModel
	var tsk model.TaskModel

	c.ShouldBindJSON(&t)

	database.Mysql.First(&tsk, t.Id)
	result := database.Mysql.Model(&tsk).Update(&model.TaskModel{
		Name:        t.Name,
		Status:      t.Status,
		Color:       t.Color,
		Tag:         t.Tag,
		Group:       t.Group,
		Description: t.Description,
	})

	if result.Error != nil {
		log.Fatal("Failed save data")
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
	})
}

func TaskDelete(c *gin.Context) {
	database.Mysql.First(&task, 2)
	database.Mysql.Model(&task).Update(model.TaskModel{})
}
