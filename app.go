package main

import "fmt"

func main() {
	price := []float64{9.99, 2.99, 4.29}
	taxRates := []float64{0.1, 0.15, 0.05}
	taxRatesWithPrice := make(map[float64][]float64)

	for _, trv := range taxRates {
		priceWithTax := make([]float64, len(price))
		for pi, pv := range price {
			//priceWithTax = append(priceWithTax, pv*trv)
			priceWithTax[pi] = pv * trv
		}
		taxRatesWithPrice[trv] = priceWithTax
	}

	fmt.Println(taxRatesWithPrice)

}
