package main

import (
	"fmt"

	"mortgage/pkg/calculator"

	"github.com/gocarina/gocsv"
)

const (
	precision = 20
)

func main() {
	mortgage := calculator.NewMortgageCalculator(300000, 625, 15 * 12)

	monthlyPaymentPlan := mortgage.CalculateMonthlyPayment(precision)

	csvContent, err := gocsv.MarshalString(&monthlyPaymentPlan)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(csvContent)
}
