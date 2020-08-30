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

func TestNextYearDivident(t *testing.T) {
	params := shareValueParams{}

	params.lastQuarterDividend, _ = decimal.NewFromString(".77")
	params.nextYearDividendIncrease, _ = decimal.NewFromString("1.05")

	expected, _ := decimal.NewFromString("3.234")
	result := params.nextYearDividend()

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestCalculateShareValue(t *testing.T) {
	params := shareValueParams{}

	params.volatility, _ = decimal.NewFromString("1.08")
	params.riskFreeRate, _ = decimal.NewFromString(".0178")
	params.historicalRisk, _ = decimal.NewFromString(".0626")
	params.lastQuarterDividend, _ = decimal.NewFromString(".77")
	params.nextYearDividendIncrease, _ = decimal.NewFromString("1.05")
	params.growthRate, _ = decimal.NewFromString("0.02")

	expected, _ := decimal.NewFromString("49.4434931506849315")
	result := params.calculateShareValue()

	if !result.Equals(expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
