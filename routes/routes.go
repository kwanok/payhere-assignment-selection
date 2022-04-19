package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/endpoints/auth"
	"github.com/payhere-assignment-selection/repository"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "We got Gin",
		})
	})

	r.GET("/users", func(c *gin.Context) {
		fmt.Println("users")
		repo := repository.UserRepository{Db: repository.DBCon}
		users := repo.GetAllUsers()
		c.JSON(200, users)
	})

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)
	r.POST("/refresh", auth.Refresh)
}
