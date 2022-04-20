package models

type Pay interface {
	GetId() string
	GetUserId() string
	GetMemo() string
	GetPrice() int
	GetRemoved() bool
	GetCreatedAt() string
}

type PayRepository interface {
	FindPayById(id string, user User) Pay
	GetPaysByUserId(userId string) []Pay
	GetRemovedPaysByUserId(userId string) []Pay
	AddPay(pay Pay, user User)
	EditPay(pay Pay)
	RemovePay(id string, user User)
	RestorePay(id string, user User)
}
