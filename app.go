package main

import "fmt"

func getName(name string) string {
	return name
}

func getAge(age int) int {
	return age
}

func double(i int) int {
	return i * 2
}

func triple(i int) int {
	return i * 3
}

// params: array, func()
func transformNumber(i []int, transform func(int) int) []int {
	var result []int

	for _, v := range i {
		result = append(result, transform(v))
	}

	return result
}

func main() {
	fmt.Println(getName("john doe"))
	fmt.Println(getAge(21))

	var numbers = []int{1, 2, 3}

	// double
	d := transformNumber(numbers, double)
	fmt.Println(d)

	// triple
	t := transformNumber(numbers, triple)
	fmt.Println(t)
}
