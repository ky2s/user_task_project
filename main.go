package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	//connect DB
	// db, err := gorm.Open(sqlite.Open("user_task_project.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// dsn := "root:@tcp(127.0.0.1:3306)/user_task_project?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "host=localhost user=postgres password=12345678 dbname=user_task_project port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect err")
	}

	//automigrate
	db.AutoMigrate(&Users{}, &Tasks{})

	// db.Migrator().AutoMigrate(&Users{})

	// // Create table for `User`
	// db.Migrator().CreateTable(&Users{})

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic("failed to connect err")

	// }

	db.Create(&Users{
		Name:  "jinzhu",
		Email: "jinzhu@gmail.com",
	})

	r := gin.Default()

	fmt.Println("connect", db)

	r.GET("/users", func(c *gin.Context) {

		// Get all records
		var users []Users
		err := db.Find(&users).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"data": nil,
				"err":  err.Error(),
			})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
		return
	})

	r.Run()
}

type Users struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tasks struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;autoIncrement"`
	UserID      int    `gorm:"index"`
	Title       string `gorm:"size:255"`
	Description string
	Status      string `gorm:"size:50"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
