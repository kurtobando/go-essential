package prices

import "fmt"

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

func (p *Price) GetPricesWithTax() {
	for _, v := range p.PricesWithTax {
		fmt.Printf("Price %.2f, Tax %.2f, Calculated Tax %.2f\n", v.Price, v.Tax, v.PriceCalculatedTax)
	}
}
