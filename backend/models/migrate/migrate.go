package main

import (
	"fmt"
	"project_management/database"
	model "project_management/models"
)

func main() {
	// START GORM DATABASE
	database.Mysql, database.Err = database.ConnectToDB()
	if database.Err != nil {
		fmt.Println("status error : ", database.Err)
	} else {
		fmt.Println("=>\t database connection \t\t✅")
	}
	defer database.Mysql.Close()
	var boxModel model.BoxModel
	var taskModel model.TaskModel
	database.Mysql.AutoMigrate(boxModel)
	database.Mysql.AutoMigrate(taskModel)
}
