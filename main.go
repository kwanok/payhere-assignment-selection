package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/config"
	"github.com/payhere-assignment-selection/repository"
	"github.com/payhere-assignment-selection/routes"
	"os"
)

func main() {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	fmt.Println(os.Getenv("DB_HOST"))

	config.InitRedis()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routes.Routes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
