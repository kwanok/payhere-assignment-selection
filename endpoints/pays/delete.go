package pays

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

func DeletePay(c *gin.Context) {
	var user, _ = GetUserInfoByToken(c)

	id := c.Param("id")

	payRepository := repository.PayRepository{Db: repository.DBCon}
	payRepository.RemovePay(id, user)

	c.JSON(http.StatusNoContent, "Successfully Deleted")
}
