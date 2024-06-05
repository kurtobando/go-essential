package main

import (
	"example.com/go-essential/user"
	"fmt"
)

func main() {
	usr, err := user.New("John Doe", 21, "New York")

	if err != nil {
		fmt.Println(err)
		return
	}

	usr.PrintUser()
	usr.SetUserName("Mary Doe")
	usr.PrintUser()

	admn, err := user.NewAdmin("admin@domain.com", "password1234!!")

	if err != nil {
		fmt.Println(err)
		return
	}

	admn.PrintUser()
	admn.SetUserName("Administrator")
	admn.PrintUser()
	admn.PrintAdmin()
}
