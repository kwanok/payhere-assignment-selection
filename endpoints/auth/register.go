package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

func Register(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	roomRepository := repository.UserRepository{Db: repository.DBCon}
	roomRepository.AddUser(&repository.User{
		Id:       uuid.New().String(),
		Email:    json.Email,
		Password: auth.Hash(json.Password),
		Name:     json.Name,
	})

	c.JSON(http.StatusOK, json)
}
