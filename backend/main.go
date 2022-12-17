package main

import (
	"fmt"
	"project_management/database"
	"project_management/router"
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
	router.Run()

}
