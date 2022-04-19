package tests

import (
	"github.com/google/uuid"
	"github.com/payhere-assignment-selection/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var pay = &repository.Pay{
	Id:     uuid.New().String(),
	UserId: user.GetId(),
	Price:  5000,
	Memo:   "인형뽑기",
}

var pays = &[]repository.Pay{
	{
		Id:     uuid.New().String(),
		UserId: user.GetId(),
		Price:  931000,
		Memo:   "아이패드 pro 11",
	},
	{
		Id:     uuid.New().String(),
		UserId: user.GetId(),
		Price:  4310000,
		Memo:   "맥북 프로 M1 Max",
	},
}

// TestAddPay 소비 내역 데이터베이스 저장하기
func TestAddPay(t *testing.T) {
	TestAddUser(t)

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	repo.AddPay(pay, user)
}

func TestAddPays(t *testing.T) {
	TestAddUser(t)

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	for _, pay := range *pays {
		repo.AddPay(&pay, user)
	}

}

// TestGetPaysByUserId 유저 아이디로 소비 내역 가져오기
func TestGetPaysByUserId(t *testing.T) {
	TestAddPays(t)

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	pays := repo.GetPaysByUserId(user.GetId())
	assert.NotEmpty(t, pays)

	for _, pay := range pays {
		assert.Equal(t, pay.GetUserId(), user.GetId())
	}
}

// TestRemovePay 소비 내역 삭제하기
func TestRemovePay(t *testing.T) {
	TestAddPays(t)

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.PayRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

}
