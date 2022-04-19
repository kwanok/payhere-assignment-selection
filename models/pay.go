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
	EditMemo(memo string)
	GetPaysByUserId(userId string) []Pay
	GetPayById(id string) Pay
	RemovePay(id string)
	AddPay(pay Pay, user User)
}
