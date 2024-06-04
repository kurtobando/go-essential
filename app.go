package main

import "fmt"

type user struct {
	name     string
	age      int
	location string
}

func main() {
	var appUser user = user{
		name:     "John Doe",
		age:      21,
		location: "New York",
	}

	fmt.Println(appUser.name)
	fmt.Println(appUser.age)
	fmt.Println(appUser.location)
}
