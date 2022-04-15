package model

type Address struct {
	State, City string
	PinCode     int
}

type User struct {
	Name    string
	Age     int
	Address Address
}
