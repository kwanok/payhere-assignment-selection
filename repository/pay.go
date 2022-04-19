package repository

import (
	"database/sql"
	"github.com/payhere-assignment-selection/models"
	"log"
)

type Pay struct {
	Id        string
	UserId    string
	Memo      string
	Removed   bool
	Price     int
	CreatedAt string
	UpdatedAt string
}

type PayRepository struct {
	Db *sql.DB
}

func (repo *PayRepository) Close() {
	repo.Db.Close()
}

func (pay *Pay) GetId() string {
	return pay.Id
}

func (pay *Pay) GetUserId() string {
	return pay.UserId
}

func (pay *Pay) GetPrice() int {
	return pay.Price
}

func (pay *Pay) GetMemo() string {
	return pay.Memo
}

func (pay *Pay) GetRemoved() bool {
	return pay.Removed
}

func (pay *Pay) GetCreatedAt() string {
	return pay.CreatedAt
}

func (repo *PayRepository) AddPay(pay models.Pay, user models.User) {
	stmt, err := repo.Db.Prepare("INSERT INTO pays(id, user_id, price, memo) values(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(pay.GetId(), user.GetId(), pay.GetPrice(), pay.GetMemo())
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *PayRepository) GetPaysByUserId(userId string) []models.Pay {
	var pays []models.Pay
	rows, err := repo.Db.Query("SELECT id, user_id, price, memo, created_at, removed FROM pays where user_id = ?", userId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var pay Pay
		rows.Scan(&pay.Id, &pay.UserId, &pay.Price, &pay.Memo, &pay.CreatedAt, &pay.Removed)
		pays = append(pays, &pay)
	}

	return pays
}

func (repo *PayRepository) EditMemo(memo string) {

}

func (repo *PayRepository) GetPayById(id string) models.Pay {
	var pay Pay

	return &pay
}

func (repo *PayRepository) RemovePay(id string) {

}
