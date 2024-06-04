package main

import "fmt"

// create type user
type user struct {
	name     string
	age      int
	location string
}

// attach a method to user struct
func (u user) displayUser() {
	fmt.Println(u.name)
	fmt.Println(u.age)
	fmt.Println(u.location)
}

func main() {
	var appUser user = user{
		name:     "John Doe",
		age:      21,
		location: "New York",
	}
	appUser.displayUser()
}
