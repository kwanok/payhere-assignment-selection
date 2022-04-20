package tests

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeUUID(t *testing.T) {
	_uuid := uuid.New()
	fmt.Println(_uuid)

	assert.NotEmpty(t, _uuid)
	assert.Equal(t, 16, len(_uuid))
}

var user = &repository.User{
	Id:       uuid.New().String(),
	Name:     "kwanok",
	Email:    "cloq@kakao.com",
	Password: auth.Hash("password"),
}

func TestGetEmail(t *testing.T) {
	assert.Equal(t, "cloq@kakao.com", user.GetEmail())
}

func TestAddUser(t *testing.T) {
	user.Id = uuid.New().String()

	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	repo.AddUser(user)
}

func TestGetAllUsers(t *testing.T) {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	users := repo.GetAllUsers()
	assert.NotNil(t, users)
}

func TestFindUserByEmail(t *testing.T) {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	dbUser := repo.FindUserByEmail(user.GetEmail())
	assert.Equal(t, user.GetEmail(), dbUser.GetEmail())
	assert.Equal(t, user.GetName(), dbUser.GetName())
}

// TestFindUserById 아이디로 유저 가져오기
func TestFindUserById(t *testing.T) {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	repo.AddUser(user)

	dbUser := repo.FindUserById(user.GetId())
	assert.Equal(t, user.GetEmail(), dbUser.GetEmail())
	assert.Equal(t, user.GetName(), dbUser.GetName())
}

// TestFindUserByIdWhenIdNull 아이디가 없을 경우
func TestFindUserByIdWhenIdNull(t *testing.T) {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	dbUser := repo.FindUserById("")
	assert.Nil(t, dbUser)
}
