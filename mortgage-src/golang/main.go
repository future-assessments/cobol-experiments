package main

import (
	"encoding/json"
	"fmt"
	"math"
)

type Customer struct {
	ID              int     `json:"id"`
	Surname         string  `json:"surname"`
	Initial         string  `json:"initial"`
	FirstName       string  `json:"firstName"`
	LoanAmount      int64   `json:"loanAmount"`
	LoanRate        float64 `json:"loanRate"`
	LoanTerm        int     `json:"loanTerm"`
	Loantype        string  `json:"loanType"`
	LoanTermMonths  int     `json:"LoanTermMonths"`
	MonthlyInterest float64 `json:"MonthlyInterest"`
	LoanRateDecimal float64 `json:"LoanRateDecimal"`
}

func NewCustomerRecord() *Customer {
	c := &Customer{
		ID:         1,
		Surname:    "Ikeda",
		Initial:    "X",
		FirstName:  "Anthony",
		LoanAmount: 10000,
		LoanRate:   3.50,
		LoanTerm:   5,
		Loantype:   "f",
	}
	c.LoanRateDecimal = c.LoanRate / 100
	c.LoanTermMonths = c.LoanTerm * 12
	c.MonthlyInterest = c.LoanRateDecimal / 12

	return c
}

func (c *Customer) Payments() {

	exponent := math.Pow(1+c.MonthlyInterest, float64(c.LoanTermMonths)) // 1 plus r to the n
	fmt.Printf("(1 + rate) ^ loanMonths: %f\n", exponent)
	numerator := c.MonthlyInterest * exponent
	fmt.Printf("rate * exponent: %f\n", numerator)

	denominator := exponent - 1
	fmt.Printf("((1 + r)^n) - 1: %f\n", denominator)

	paymentMultiplier := numerator / denominator
	fmt.Printf("payment multiplier: %f\n", paymentMultiplier)
	payment := float64(c.LoanAmount) * paymentMultiplier
	fmt.Printf("payment: %f\n", payment)

}

func main() {
	c := NewCustomerRecord()
	ob, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(ob))
	c.Payments()
}
