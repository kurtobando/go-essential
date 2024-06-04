package main

import (
	"errors"
	"fmt"
)

// create type user
type user struct {
	name     string
	age      int
	location string
}

// attach a method to user struct
func (u *user) displayUser() {
	fmt.Printf("%s %d %s \n", u.name, u.age, u.location)
}

// attach a method that will mutate the params
func (u *user) setName(name string) {
	u.name = name
}

// convention when creating a struct
func newUser(name string, age int, location string) (*user, error) {
	if name == "" || age == 0 || location == "" {
		return nil, errors.New("invalid arguments")
	}
	return &user{
		name:     name,
		age:      age,
		location: location,
	}, nil
}

func main() {
	_user, err := newUser("John Doe", 21, "New York")
	if err != nil {
		fmt.Println(err)
	}
	_user.displayUser()
	_user.setName("Mary Doe")
	_user.displayUser()

	// return memory address
	fmt.Println(&_user)
	// return the value of the memory address
	fmt.Println(*_user)
}
