package model

type Address struct {
	State, City string
	PinCode     int
}

type User struct {
	Name    string `json:"name" bson:"user_name"`
	Age     int
	Address Address
}
