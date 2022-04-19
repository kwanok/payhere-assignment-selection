package models

type User interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string
}

type UserRepository interface {
	AddUser(user User)
	RemoveUser(user User)
	FindUserById(Id string) User
	GetAllUsers() []User
}
