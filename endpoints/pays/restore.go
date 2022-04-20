package pays

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

// RestorePay 삭제된 내역 복구
func RestorePay(c *gin.Context) {
	var user, _ = GetUserInfoByToken(c)

	id := c.Param("id")

	payRepository := &repository.PayRepository{Db: repository.DBCon}

	payRepository.RestorePay(id, user)

	c.JSON(http.StatusOK, "")
}
