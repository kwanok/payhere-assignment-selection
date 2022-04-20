package repository

import (
	"database/sql"
	"github.com/payhere-assignment-selection/models"
	"log"
)

type Pay struct {
	Id        string `json:"id"`
	UserId    string `json:"-"`
	Memo      string `json:"memo"`
	Removed   bool   `json:"-"`
	Price     int    `json:"price"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
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
	rows, err := repo.Db.Query("SELECT id, user_id, price, memo, created_at, updated_at, removed FROM pays WHERE user_id = ? AND removed=0", userId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var pay Pay
		rows.Scan(&pay.Id, &pay.UserId, &pay.Price, &pay.Memo, &pay.CreatedAt, &pay.UpdatedAt, &pay.Removed)
		pays = append(pays, &pay)
	}

	return pays
}

func (repo *PayRepository) GetRemovedPaysByUserId(userId string) []models.Pay {
	var pays []models.Pay
	rows, err := repo.Db.Query("SELECT id, user_id, price, memo, created_at, updated_at, removed FROM pays WHERE user_id = ? AND removed=1", userId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var pay Pay
		rows.Scan(&pay.Id, &pay.UserId, &pay.Price, &pay.Memo, &pay.CreatedAt, &pay.UpdatedAt, &pay.Removed)
		pays = append(pays, &pay)
	}

	return pays
}

func (repo *PayRepository) FindPayById(id string, user models.User) models.Pay {
	row := repo.Db.QueryRow("SELECT id, user_id, price, memo, created_at, updated_at, removed FROM pays where id = ? AND user_id = ? AND removed=0 LIMIT 1", id, user.GetId())

	var pay Pay

	if err := row.Scan(&pay.Id, &pay.UserId, &pay.Price, &pay.Memo, &pay.CreatedAt, &pay.UpdatedAt, &pay.Removed); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err)
	}

	return &pay
}

func (repo *PayRepository) EditPay(pay models.Pay) {
	stmt, err := repo.Db.Prepare("UPDATE pays SET price=?,memo=? WHERE id=? AND removed=0")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(pay.GetPrice(), pay.GetMemo(), pay.GetId())
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *PayRepository) RemovePay(id string, user models.User) {
	stmt, err := repo.Db.Prepare("UPDATE pays SET removed=1 WHERE id=? AND user_id=? AND removed=0")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id, user.GetId())
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *PayRepository) RestorePay(id string, user models.User) {
	stmt, err := repo.Db.Prepare("UPDATE pays SET removed=0 WHERE id=? AND user_id=? AND removed=1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id, user.GetId())
	if err != nil {
		log.Fatal(err)
	}
}
