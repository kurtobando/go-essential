package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var location string

	// Reads input from the standard input, stops reading at a newline, and discards the remaining characters on the line.
	// Useful when you want to read input until the end of the line.
	fmt.Print("What is your name? ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Printf("Hello %s!\n", name)

	// Reads formatted input from the standard input based on a format string.
	// Useful when you need to read input in a specific format.
	fmt.Print("What is your age? ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		return
	}
	fmt.Printf("Your age is: %d\n", age)

	// Reads input from the standard input and stops reading at whitespace.
	// Useful for reading space-separated values.
	fmt.Print("What is your location? ")
	_, err = fmt.Scan(&location)
	if err != nil {
		return
	}
	fmt.Printf("You are located at %s.\n", location)
	fmt.Printf("Type %T\n", location)

	// fmt formats
	// %s string
	// %d decimal integer
	// %f floating-point number.
	// %v boolean
	// %T type of the value
}
