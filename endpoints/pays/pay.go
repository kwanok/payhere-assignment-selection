package pays

import (
	"github.com/gin-gonic/gin"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/models"
	"github.com/payhere-assignment-selection/repository"
	"net/http"
)

type Pay struct {
	Id    string `json:"id"`
	Price string `json:"price"`
	Memo  string `json:"memo"`
}

func GetUserInfoByToken(c *gin.Context) (models.User, error) {
	// 레디스 저장소에서 유저 ID가 담긴 토큰 가져오기
	accessDetail, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		c.Abort()
	}

	userId, err := auth.FetchAuth(accessDetail)

	userRepository := &repository.UserRepository{Db: repository.DBCon}

	user := userRepository.FindUserById(userId)
	if user == nil {
		c.JSON(http.StatusUnauthorized, "unauthorized - user not exists")
		c.Abort()
	}

	return user, nil
}
