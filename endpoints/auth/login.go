package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

func Login(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userRepository := repository.UserRepository{Db: repository.DBCon}
	user := userRepository.FindUserByEmail(json.Email)

	if user == nil {
		c.JSON(http.StatusNotFound, "Not found")
		c.Abort()
		return
	}

	if user.GetEmail() != json.Email || !auth.CompareHash(user.GetPassword(), json.Password) {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := auth.CreateToken(user.GetId())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "StatusUnprocessableEntity")
		return
	}

	saveErr := auth.CreateAuth(user.GetId(), token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}
