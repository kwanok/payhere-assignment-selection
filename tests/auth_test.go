package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/payhere-assignment-selection/config"
	"github.com/payhere-assignment-selection/config/auth"
	auth2 "github.com/payhere-assignment-selection/endpoints/auth"
	"github.com/payhere-assignment-selection/repository"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

// TestRegister 고객은 이메일과 비밀번호 입력을 통해서 회원 가입을 할 수 있습니다.
func TestRegister(t *testing.T) {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	rPath := "/register"
	router := gin.Default()
	router.POST(rPath, auth2.Register)
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{"email":"test@example.com", "nickname":"knoh", "password":"password!"}`))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

// TestLogin 고객은 회원 가입이후, 로그인을 할 수 있습니다.
func TestLogin(t *testing.T) {
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

	rPath := "/login"
	router := gin.Default()
	router.POST(rPath, auth2.Login)
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{"email":"rosejap97@gmail.com", "password":"password!"}`))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestRefresh(t *testing.T) {
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

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	repo.AddUser(user)

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)

	rPath := "/refresh"
	router := gin.Default()
	router.POST(rPath, auth2.Refresh)
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{"refresh_token":"`+token.RefreshToken+`"}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 201, w.Code)
}

// TestLogout 고객은 회원 가입이후, 로그아웃을 할 수 있습니다.
func TestLogout(t *testing.T) {
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

	rPath := "/logout"
	router := gin.Default()
	router.POST(rPath, auth2.Logout)
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}
