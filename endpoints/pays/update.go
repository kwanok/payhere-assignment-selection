package pays

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

type UpdateRequest struct {
	Price int    `json:"price"`
	Memo  string `json:"memo"`
}

func UpdatePay(c *gin.Context) {
	var user, _ = GetUserInfoByToken(c)

	id := c.Param("id")

	var json UpdateRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	payRepository := &repository.PayRepository{Db: repository.DBCon}

	pay := payRepository.FindPayById(id, user)
	if pay == nil {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	var payUpdate = &repository.Pay{
		Id:    pay.GetId(),
		Price: pay.GetPrice(),
		Memo:  pay.GetMemo(),
	}

	if pay.GetPrice() != json.Price {
		payUpdate.Price = json.Price
	}
	if pay.GetMemo() != json.Memo {
		payUpdate.Memo = json.Memo
	}

	payRepository.EditPay(payUpdate)

	c.JSON(http.StatusNoContent, "Successfully Updated")
}
