package main

import (
	"fmt"
	"math"

	"github.com/gocarina/gocsv"
)

type Customer struct {
	ID               int        `json:"id" csv:"id"`
	Surname          string     `json:"surname" csv:"surname"`
	Initial          string     `json:"initial" csv:"initial"`
	FirstName        string     `json:"firstName" csv:"first_name"`
	LoanAmount       int64      `json:"loanAmount" csv:"loan_amount"`
	LoanRate         float64    `json:"loanRate" csv:"loan_rate"`
	LoanTerm         int        `json:"loanTerm" csv:"loan_term"`
	Loantype         string     `json:"loanType" csv:"loan_type"`
	LoanTermMonths   int        `json:"LoanTermMonths" csv:"loan_term_months"`
	MonthlyInterest  float64    `json:"MonthlyInterest" csv:"monthly_interest"`
	LoanRateDecimal  float64    `json:"LoanRateDecimal" csv:"loan_rate_decimal"`
	PaymentsSchedule []*Payment `json:"payment_schedule"`
	MonthlyPayment   float64    `json:"monthly_payment"`
}

// top -> down order of struct fields correlates to left->right order of CSV fields.
type Payment struct {
	PaymentNumber int     `json:"payment_number" csv:"payment_number"`
	Principal     float64 `json:"principal" csv:"principal"`
	Interest      float64 `json:"interest" csv:"interest"`
	Balance       float64 `json:"balance" csv:"balance"`
}

func NewCustomerRecord() *Customer {
	c := &Customer{
		ID:         1,
		Surname:    "Ikeda",
		Initial:    "X",
		FirstName:  "Anthony",
		LoanAmount: 300000,
		LoanRate:   6.25,
		LoanTerm:   15,
		Loantype:   "f",
	}
	c.LoanRateDecimal = c.LoanRate / 100
	c.LoanTermMonths = c.LoanTerm * 12
	c.MonthlyInterest = c.LoanRateDecimal / 12
	c.calculateMonthlyPayment()
	return c
}

func (c *Customer) Payments() {

	//payment 1
	c.calculatePayment(float64(c.LoanAmount))

	for i := 2; i <= c.LoanTermMonths; i++ {
		bal := c.PaymentsSchedule[(len(c.PaymentsSchedule) - 1)].Balance
		c.calculatePayment(bal)
	}
}

func (c *Customer) calculateMonthlyPayment() {
	exponent := math.Pow(1+c.MonthlyInterest, float64(c.LoanTermMonths)) // 1 plus r to the n
	//fmt.Printf("(1 + rate) ^ loanMonths: %f\n", exponent)
	numerator := c.MonthlyInterest * exponent
	//fmt.Printf("rate * exponent: %f\n", numerator)

	denominator := exponent - 1
	//fmt.Printf("((1 + r)^n) - 1: %f\n", denominator)

	paymentMultiplier := numerator / denominator
	//fmt.Printf("payment multiplier: %f\n", paymentMultiplier)
	payment := float64(c.LoanAmount) * paymentMultiplier
	c.MonthlyPayment = payment
}

func (c *Customer) calculatePayment(balance float64) {

	interestPayment := float64(balance) * c.MonthlyInterest

	p := &Payment{
		Interest:      float64(balance) * c.MonthlyInterest,
		Principal:     c.MonthlyPayment - interestPayment,
		PaymentNumber: len(c.PaymentsSchedule) + 1,
	}
	p.Balance = float64(balance) - p.Principal
	c.PaymentsSchedule = append(c.PaymentsSchedule, p)

}

func PCSV(d any) {
	csvContent, err := gocsv.MarshalString(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(csvContent)
}
func main() {
	c := NewCustomerRecord()
	c.Payments()
	PCSV(c.PaymentsSchedule)
}
