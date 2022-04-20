package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/payhere-assignment-selection/config"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/endpoints/pays"
	"github.com/payhere-assignment-selection/middlewares"
	"github.com/payhere-assignment-selection/repository"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func newPay() *repository.Pay {
	return &repository.Pay{
		Id:    uuid.New().String(),
		Price: 54000,
		Memo:  "메모다 메모!",
	}
}

// TestStorePay 가계부에 오늘 사용한 돈의 금액과 관련된 메모를 남길 수 있습니다.
func TestStorePay(t *testing.T) {
	user := newUser()

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

	rPath := "/pays"
	router := gin.Default()
	router.POST(rPath, middlewares.IsAuthorized, pays.CreatePay)
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{"price":30000, "memo":"메모~~"}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 201, w.Code)
}

// TestEditPay 가계부에서 수정을 원하는 내역은 금액과 메모를 수정 할 수 있습니다.
func TestEditPay(t *testing.T) {
	user := newUser()
	pay := newPay()

	config.InitRedis()

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	userRepo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		userRepo.Close()
	}()

	// 유저 생성
	userRepo.AddUser(user)

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)

	// 소비 내역 생성
	payRepo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		payRepo.Close()
	}()

	payRepo.AddPay(pay, user)

	rPath := "/pays/:id"
	router := gin.Default()
	router.PATCH(rPath, middlewares.IsAuthorized, pays.UpdatePay)
	req, _ := http.NewRequest("PATCH", "/pays/"+pay.GetId(), strings.NewReader(`{"price":80000, "memo":"메모~~ 수정.."}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 204, w.Code)

	assert.Equal(t, "메모~~ 수정..", payRepo.FindPayById(pay.Id, user).GetMemo())
	assert.Equal(t, 80000, payRepo.FindPayById(pay.Id, user).GetPrice())
}

// TestDeletePay 가계부에서 삭제를 원하는 내역은 삭제 할 수 있습니다.
func TestDeletePay(t *testing.T) {
	user := newUser()
	pay := newPay()

	config.InitRedis()

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	userRepo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		userRepo.Close()
	}()

	// 유저 생성
	userRepo.AddUser(user)

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)

	// 소비 내역 생성
	payRepo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		payRepo.Close()
	}()

	payRepo.AddPay(pay, user)

	rPath := "/pays/:id"
	router := gin.Default()
	router.DELETE(rPath, middlewares.IsAuthorized, pays.DeletePay)
	req, _ := http.NewRequest("DELETE", "/pays/"+pay.GetId(), strings.NewReader(`{""}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, nil, payRepo.FindPayById(pay.Id, user))
}

// TestRestorePay 삭제한 내역은 언제든지 다시 복구 할 수 있어야 한다.
func TestRestorePay(t *testing.T) {
	user := newUser()
	pay := newPay()

	config.InitRedis()

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	userRepo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		userRepo.Close()
	}()

	// 유저 생성
	userRepo.AddUser(user)

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)

	// 소비 내역 생성
	payRepo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		payRepo.Close()
	}()

	payRepo.AddPay(pay, user)
	payRepo.RemovePay(pay.GetId(), user)

	rPath := "/pays/restore/:id"
	router := gin.Default()
	router.PATCH(rPath, middlewares.IsAuthorized, pays.RestorePay)
	req, _ := http.NewRequest("PATCH", "/pays/restore/"+pay.GetId(), strings.NewReader(`{""}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, payRepo.FindPayById(pay.GetId(), user).GetRemoved())
}

// TestGetPays 가계부에서 이제까지 기록한 가계부 리스트를 볼 수 있습니다.
func TestGetPays(t *testing.T) {
	user := newUser()

	config.InitRedis()
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	userRepo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		userRepo.Close()
	}()

	// 유저 생성
	userRepo.AddUser(user)

	payRepo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		payRepo.Close()
	}()

	for i := 1; i <= 5; i++ {
		payRepo.AddPay(&repository.Pay{
			Id:    uuid.New().String(),
			Price: i * 10000,
			Memo:  strconv.Itoa(i) + "번째 메모",
		}, user)
	}

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)

	rPath := "/pays"
	router := gin.Default()
	router.GET(rPath, middlewares.IsAuthorized, pays.GetPays)
	req, _ := http.NewRequest("GET", "/pays", strings.NewReader(`{""}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)

	var result []repository.Pay

	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 5개가 잘 뽑혔는지 확인
	assert.Equal(t, 5, len(result))
}

// TestGetPay 가계부에서 상세한 세부 내역을 볼 수 있습니다.
func TestGetPay(t *testing.T) {
	pay := newPay()
	user := newUser()

	config.InitRedis()
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	userRepo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		userRepo.Close()
	}()

	userRepo.AddUser(user)

	token, _ := auth.CreateToken(user.GetId())
	saveErr := auth.CreateAuth(user.GetId(), token)
	assert.Nil(t, saveErr)

	payRepo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		payRepo.Close()
	}()

	payRepo.AddPay(pay, user)

	rPath := "/pays/:id"
	router := gin.Default()
	router.GET(rPath, middlewares.IsAuthorized, pays.GetPay)
	req, _ := http.NewRequest("GET", "/pays/"+pay.GetId(), strings.NewReader(`{""}`))
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)

	var result repository.Pay

	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		log.Fatalln(err)
		return
	}

	assert.Equal(t, pay.GetPrice(), result.GetPrice())
	assert.Equal(t, pay.GetId(), result.GetId())
	assert.Equal(t, pay.GetMemo(), result.GetMemo())
}
