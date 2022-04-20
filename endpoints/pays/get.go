package pays

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/models"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

// GetPays
func GetPays(c *gin.Context) {
	var user, _ = GetUserInfoByToken(c)

	var pays []models.Pay

	payRepository := repository.PayRepository{Db: repository.DBCon}
	pays = payRepository.GetPaysByUserId(user.GetId())

	c.JSON(http.StatusOK, pays)
}

// GetPay
func GetPay(c *gin.Context) {
	var user, _ = GetUserInfoByToken(c)

	id := c.Param("id")

	var pay models.Pay

	payRepository := repository.PayRepository{Db: repository.DBCon}
	pay = payRepository.FindPayById(id, user)
	if pay == nil {
		c.JSON(http.StatusNotFound, "Not Found")
	}

	c.JSON(http.StatusOK, pay)
}

// getRemovedPays
func getRemovedPays(c *gin.Context) {

}
