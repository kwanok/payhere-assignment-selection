package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/endpoints/auth"
	"github.com/payhere-assignment-selection/endpoints/pays"
	"github.com/payhere-assignment-selection/middlewares"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "We got Gin",
		})
	})

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/logout", middlewares.IsAuthorized, auth.Logout)
	r.POST("/refresh", auth.Refresh)

	payGroup := r.Group("/pays", middlewares.IsAuthorized)
	{
		payGroup.GET("/", pays.GetPays)
		payGroup.GET("/:id", pays.GetPay)
		payGroup.POST("/", pays.CreatePay)
		payGroup.PATCH("/:id", pays.UpdatePay)
		payGroup.DELETE("/:id", pays.DeletePay)
		payGroup.PATCH("/:id", pays.RestorePay)
	}
}
