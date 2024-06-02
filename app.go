package main

import "fmt"

func main() {
	var priceItem1 = 100
	var priceItem2 = 200
	fmt.Println("Result: ", priceItem1+priceItem2)

	var priceItem3 = 312.08
	var priceItem4 float64 = 400
	fmt.Println("Result: ", priceItem3+priceItem4)

	var f32 float32 = 3.141592653589793238
	var f64 float64 = 3.141592653589793238
	fmt.Printf("float32: %.20f\n", f32)
	fmt.Printf("float64: %.20f\n", f64)
}
