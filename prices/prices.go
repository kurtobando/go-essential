package prices

import (
	"example.com/go-essential/filemanager"
	"example.com/go-essential/helper"
	"fmt"
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
	readFile, err := filemanager.ReadFile("tmp/prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	float, err := helper.StringsToFloat(readFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	p.Prices = float
}

func (p *Price) GetPricesWithTax() {
	fmt.Println("***********************")
	for _, v := range p.PricesWithTax {
		fmt.Printf("Price %.2f, Tax %.2f, Calculated Tax %.2f\n", v.Price, v.Tax, v.PriceCalculatedTax)
		_ = filemanager.WriteFile(fmt.Sprintf("tmp/prices-with-tax-%.0f.json", v.Price), p)
	}

}
