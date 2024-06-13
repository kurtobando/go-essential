package main

import "example.com/go-essential/prices"

func main() {
	productPrices := []float64{9.99, 2.99, 4.29}
	taxRates := []float64{0.1, 0.15, 0.05}

	for _, v := range taxRates {
		p := prices.New(v, productPrices)
		p.Process()
		p.GetPricesWithTax()
	}
}
