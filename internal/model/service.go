package model

type Service interface {
	GetHomePage() string
	AddUser(user User) error
}
