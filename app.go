package main

import (
	"errors"
	"fmt"
)

type User struct {
	name     string
	age      int
	location string
}

// attach a method to User struct
func (u *User) displayUser() {
	fmt.Printf("%s %d %s \n", u.name, u.age, u.location)
}

// attach a method that will mutate the params
func (u *User) setName(name string) {
	u.name = name
}

// convention when creating a struct
func newUser(name string, age int, location string) (*User, error) {
	if name == "" || age == 0 || location == "" {
		return nil, errors.New("invalid arguments")
	}
	return &User{
		name:     name,
		age:      age,
		location: location,
	}, nil
}

func main() {
	user, err := newUser("John Doe", 21, "New York")
	if err != nil {
		fmt.Println(err)
	}
	user.displayUser()
	user.setName("Mary Doe")
	user.displayUser()

	// return memory address
	fmt.Println(&user)
	// return the value of the memory address
	fmt.Println(*user)
}
