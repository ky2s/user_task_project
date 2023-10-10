package main

import (
	"fmt"
	"user_task_project/controllers"
	"user_task_project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=12345678 dbname=user_task_project port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect err")
	}

	//automigrate
	db.AutoMigrate(&models.Users{})

	// db.Create(&Users{
	// 	Name:  "jinzhu",
	// 	Email: "jinzhu@gmail.com",
	// })
	var userModels models.UserModels = models.NewUserModels(db)
	var userController controllers.UserController = controllers.NewUserController(userModels)

	r := gin.Default()

	fmt.Println("connect", db)

	// endpoint user
	r.POST("/users", userController.InsertUser)        //done
	r.GET("/users", userController.GetUser)            //done
	r.GET("/users/:id", userController.GetUser)        //done
	r.PUT("/users/:id", userController.UpdateUser)     //done
	r.DELETE("/users/:id", userController.DestroyUser) //done

	// endpoint tasks
	r.POST("/tasks", userController.GetUser)
	r.GET("/tasks", userController.GetUser)
	r.GET("/tasks/:id", userController.GetUser)
	r.PUT("/tasks/:id", userController.GetUser)
	r.DELETE("/tasks/:id", userController.GetUser)

	r.Run()
}
