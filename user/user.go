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
