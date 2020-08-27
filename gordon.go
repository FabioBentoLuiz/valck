// Gordon Growth Model implementation.
// This model works well only for companies that derive theirs primary value from dividends (quite a few if any).
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type shareValueParams struct {
	nextYearDividend decimal.Decimal
	growthRate       decimal.Decimal
	riskFreeRate     decimal.Decimal
	volatility       decimal.Decimal
	historicalRisk   decimal.Decimal
}

func (p shareValueParams) calculate() decimal.Decimal {
	shareValue := p.nextYearDividend.Div(p.discountRate().Sub(p.growthRate))

	return shareValue
}

func (p shareValueParams) discountRate() decimal.Decimal {
	discountRate := p.riskFreeRate.Add(p.volatility.Mul(p.historicalRisk))
	return discountRate
}

func (p shareValueParams) test() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
