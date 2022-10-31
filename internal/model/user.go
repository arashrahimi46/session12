package model

type User struct {
	Id       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Age      int    `json:"age" bson:"age"`
	MobileNo string `json:"mobile_no" bson:"mobile_no"`
}
