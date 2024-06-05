package user

import (
	"errors"
	"fmt"
)

type User struct {
	name     string
	age      int
	location string
}

type Admin struct {
	email    string
	password string
	User
}

func (a *Admin) PrintAdmin() {
	fmt.Printf("You are an (admin) name: %s, email: %s\n", a.name, a.email)
}

func (u *User) PrintUser() {
	fmt.Printf("%s %d %s \n", u.name, u.age, u.location)
}

func (u *User) SetUserName(name string) {
	u.name = name
}

func New(name string, age int, location string) (*User, error) {
	if name == "" || age == 0 || location == "" {
		return nil, errors.New("invalid arguments")
	}
	return &User{
		name:     name,
		age:      age,
		location: location,
	}, nil
}

func NewAdmin(email string, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("invalid arguments")
	}
	return &Admin{
		email:    email,
		password: password,
		User: User{
			name:     "administrator",
			age:      99,
			location: "IT Support",
		},
	}, nil
}
