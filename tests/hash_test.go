package tests

import (
	"fmt"
	"github.com/payhere-assignment-selection/config/auth"
	"github.com/payhere-assignment-selection/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareHash(t *testing.T) {
	repository.InitDB()
	defer func() {
		repository.DBCon.Close()
	}()

	repo := &repository.UserRepository{Db: repository.DBCon}
	defer func() {
		repo.Close()
	}()

	dbUser := repo.FindUserByEmail(user.GetEmail())
	fmt.Println(dbUser.GetEmail(), dbUser.GetPassword())
	assert.True(t, auth.CompareHash(dbUser.GetPassword(), "password"))
}
