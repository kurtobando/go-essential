package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Price struct {
	TaxRate       float64
	Prices        []float64
	PricesWithTax []PriceWithTaxCalculation
}

type PriceWithTaxCalculation struct {
	Price              float64
	PriceCalculatedTax float64
	Tax                float64
}

func New(taxRate float64, prices []float64) *Price {
	return &Price{
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (p *Price) Process() {
	for _, v := range p.Prices {
		calculation := PriceWithTaxCalculation{
			Price:              v,
			Tax:                p.TaxRate,
			PriceCalculatedTax: v * p.TaxRate,
		}
		p.PricesWithTax = append(p.PricesWithTax, calculation)
	}
}

func (p *Price) GetPricesFromAFile() {
	// open file via os library
	file, err := os.Open("tmp/prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// read file line by line
	scanner := bufio.NewScanner(file)
	var scannedLines []string
	for scanner.Scan() {
		scannedLines = append(scannedLines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert the strings from the file to float
	var prices []float64
	for _, line := range scannedLines {
		p, _ := strconv.ParseFloat(line, 64)
		prices = append(prices, p)
	}
	p.Prices = prices

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *Price) GetPricesWithTax() {
	for _, v := range p.PricesWithTax {
		fmt.Printf("Price %.2f, Tax %.2f, Calculated Tax %.2f\n", v.Price, v.Tax, v.PriceCalculatedTax)
	}
}
