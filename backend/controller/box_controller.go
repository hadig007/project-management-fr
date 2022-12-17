package controller

import (
	"log"
	"project_management/database"
	model "project_management/models"

	"github.com/gin-gonic/gin"
)

var box model.TaskModel

func BoxStore(c *gin.Context) {
	result := database.Mysql.Save(&model.BoxModel{
		Name:  "Starter",
		Color: "blue",
	})
	if result.Error != nil {
		log.Fatal("Failed save data")
	}
}

func BoxIndex(c *gin.Context) {
	data := database.Mysql.Find(&task)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}

func BoxUpdate(c *gin.Context) {
	database.Mysql.First(&box, 2)
	database.Mysql.Model(&task).Update(model.TaskModel{})
}

func BoxDelete(c *gin.Context) {
	database.Mysql.First(&box, 2)
	database.Mysql.Model(&task).Update(model.TaskModel{})
}
