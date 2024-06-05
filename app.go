package main

import (
	"example.com/go-essential/user"
	"fmt"
)

func main() {
	usr, err := user.New("John Doe", 21, "New York")

	if err != nil {
		fmt.Println(err)
	}

	usr.PrintUser()
	usr.SetUserName("Mary Doe")
	usr.PrintUser()
}
