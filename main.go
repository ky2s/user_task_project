package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	fmt.Println("connect")

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "hello",
		})
	})

	r.Run()
}
