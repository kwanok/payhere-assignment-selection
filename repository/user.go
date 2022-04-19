package repository

import (
	"database/sql"
	"github.com/payhere-assignment-selection/models"
	"log"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRepository struct {
	Db *sql.DB
}

func (repo *UserRepository) Close() {
	repo.Db.Close()
}

func (user *User) GetId() string {
	return user.Id
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) GetName() string {
	return user.Name
}

func (repo *UserRepository) AddUser(user models.User) {
	stmt, err := repo.Db.Prepare("INSERT INTO users(id, email, password, name) values(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.GetId(), user.GetEmail(), user.GetPassword(), user.GetName())
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *UserRepository) GetAllUsers() []models.User {
	rows, err := repo.Db.Query("SELECT id, email FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email)
		users = append(users, &user)
	}

	return users
}

func (repo *UserRepository) FindUserById(Id string) models.User {
	row := repo.Db.QueryRow("SELECT id, email, password, name FROM users where id = ? LIMIT 1", Id)

	var user User

	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err)
	}

	return &user
}

func (repo *UserRepository) FindUserByEmail(Email string) models.User {
	row := repo.Db.QueryRow("SELECT id, email, password, name FROM users where email = ? LIMIT 1", Email)

	var user User

	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err)
	}

	return &user
}

func (repo *UserRepository) FindUserByPay(pay Pay) models.User {
	row := repo.Db.QueryRow("SELECT id, email, password, name FROM users where id = ? LIMIT 1", pay.GetUserId())

	var user User

	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Fatal(err)
	}

	return &user
}

func (repo *UserRepository) RemoveUser(user models.User) {

}
