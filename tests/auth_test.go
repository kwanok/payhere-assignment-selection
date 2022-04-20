package tests

import (
	"github.com/joho/godotenv"
	"github.com/payhere-assignment-selection/config"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/repository"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

// TestAddToken 토큰 생성해서 저장하기
func TestAddToken(t *testing.T) {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	config.InitRedis()

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)
}
