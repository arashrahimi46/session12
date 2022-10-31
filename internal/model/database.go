package model

type Database interface {
	//InsertUser this is a function for insert new user to database
	InsertUser(user User) (lastInsertedId string, err error)
	//DeleteUser this a function for deleting user
	DeleteUser()
	//GetUser this a function for get user with id
	GetUser()
}
