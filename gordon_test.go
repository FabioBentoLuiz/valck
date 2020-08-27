package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestDiscountRate(t *testing.T) {
	params := shareValueParams{}

	params.volatility, _ = decimal.NewFromString("1.08")
	params.riskFreeRate, _ = decimal.NewFromString(".0178")
	params.historicalRisk, _ = decimal.NewFromString(".0626")

	expected, _ := decimal.NewFromString(".085408")
	result := params.discountRate()

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
