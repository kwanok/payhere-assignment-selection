package pays

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

type CreateRequest struct {
	Price int    `json:"price"`
	Memo  string `json:"memo"`
}

func CreatePay(c *gin.Context) {
	var user, _ = GetUserInfoByToken(c)

	var json CreateRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := uuid.New().String()
	payRepository := repository.PayRepository{Db: repository.DBCon}
	payRepository.AddPay(&repository.Pay{
		Id:     id,
		UserId: user.GetId(),
		Price:  json.Price,
		Memo:   json.Memo,
	}, user)

	c.JSON(http.StatusCreated, id)
}
