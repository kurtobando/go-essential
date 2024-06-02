package main

import (
	"example.com/go-essential/greetings"
	"fmt"
)

func main() {
	fmt.Println("Running ...")

	greetings.SayHelloWorld()
	greetings.SayHelloUser("John Doe")
}
