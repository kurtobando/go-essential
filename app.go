package main

import (
	"example.com/go-essential/input"
	"example.com/go-essential/note"
	"fmt"
)

func main() {
	t, err := input.GetUserInput("Title: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := input.GetUserInput("Content: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	n := note.New(t, c)
	n.PrintNote()
	n.SaveAsJson()
}
