package main

import "fmt"

func main() {
	fmt.Println(getName("john doe"))
	fmt.Println(getAge(21))
}
func getName(name string) string {
	return name
}
func getAge(age int) int {
	return age
}
