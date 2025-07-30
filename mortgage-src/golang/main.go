package main

import (
	"fmt"
	"math/big"

	"github.com/ALTree/bigfloat"
	"github.com/gocarina/gocsv"
)

const (
	FP_PRECISION = 100
)

type Customer struct {
	ID               int        `json:"id" csv:"id"`
	Surname          string     `json:"surname" csv:"surname"`
	Initial          string     `json:"initial" csv:"initial"`
	FirstName        string     `json:"firstName" csv:"first_name"`
	LoanAmount       *big.Float `json:"loanAmount" csv:"loan_amount"`
	LoanRate         *big.Float `json:"loanRate" csv:"loan_rate"`
	LoanTerm         int        `json:"loanTerm" csv:"loan_term"`
	Loantype         string     `json:"loanType" csv:"loan_type"`
	LoanTermMonths   int        `json:"LoanTermMonths" csv:"loan_term_months"`
	MonthlyInterest  *big.Float `json:"MonthlyInterest" csv:"monthly_interest"`
	LoanRateDecimal  *big.Float `json:"LoanRateDecimal" csv:"loan_rate_decimal"`
	PaymentsSchedule []*Payment `json:"payment_schedule"`
	MonthlyPayment   *big.Float `json:"monthly_payment"`
}

// top -> down order of struct fields correlates to left->right order of CSV fields.
type Payment struct {
	PaymentNumber int        `json:"payment_number" csv:"payment_number"`
	Principal     *big.Float `json:"principal" csv:"principal"`
	Interest      *big.Float `json:"interest" csv:"interest"`
	Balance       *big.Float `json:"balance" csv:"balance"`
}

func NewPayment(precision uint) *Payment {

	p := &Payment{}
	p.Principal = new(big.Float).SetPrec(precision)
	p.Interest = new(big.Float).SetPrec(precision)
	p.Balance = new(big.Float).SetPrec(precision)
	return p
}

func NewCustomerRecord(precision uint) *Customer {
	c := &Customer{
		ID:              1,
		Surname:         "Ikeda",
		Initial:         "X",
		FirstName:       "Anthony",
		LoanAmount:      new(big.Float).SetPrec(precision).SetInt64(300000),
		LoanRate:        new(big.Float).SetPrec(precision).SetFloat64(6.25),
		LoanTerm:        15,
		Loantype:        "f",
		LoanRateDecimal: new(big.Float).SetPrec(precision),
		MonthlyInterest: new(big.Float).SetPrec(precision),
	}
	c.LoanRateDecimal.Quo(c.LoanRate, big.NewFloat(100))
	c.LoanTermMonths = c.LoanTerm * 12
	c.MonthlyInterest.Quo(c.LoanRateDecimal, big.NewFloat(12))
	c.calculateMonthlyPayment()
	return c
}

func (c *Customer) Payments() {

	//payment 1
	c.calculatePayment(c.LoanAmount)

	for i := 2; i <= c.LoanTermMonths; i++ {
		bal := c.PaymentsSchedule[(len(c.PaymentsSchedule) - 1)].Balance
		c.calculatePayment(bal)
	}
}

func (c *Customer) calculateMonthlyPayment() {
	onePlusMonthlyInterest := new(big.Float).SetPrec(FP_PRECISION)
	onePlusMonthlyInterest.Add(c.MonthlyInterest, big.NewFloat(1))
	exponent := bigfloat.Pow(onePlusMonthlyInterest, big.NewFloat(float64(c.LoanTermMonths))) // 1 plus r to the n
	//fmt.Printf("(1 + rate) ^ loanMonths: %f\n", exponent)
	// numerator := c.MonthlyInterest * exponent
	numerator := new(big.Float).SetPrec(FP_PRECISION)
	numerator.Mul(c.MonthlyInterest, exponent)
	//fmt.Printf("rate * exponent: %f\n", numerator)

	// denominator := exponent - 1
	denominator := new(big.Float).SetPrec(FP_PRECISION)
	denominator.Sub(exponent, big.NewFloat(1))
	//fmt.Printf("((1 + r)^n) - 1: %f\n", denominator)

	paymentRatio := new(big.Float).SetPrec(FP_PRECISION)
	paymentRatio.Quo(numerator, denominator)
	//fmt.Printf("payment multiplier: %f\n", paymentMultiplier)
	payment := new(big.Float).SetPrec(FP_PRECISION)
	payment.Mul(paymentRatio, c.LoanAmount)
	c.MonthlyPayment = payment
	fmt.Println(c.MonthlyPayment)
}

func (c *Customer) calculatePayment(balance *big.Float) {

	interestPayment := new(big.Float).SetPrec(FP_PRECISION)
	interestPayment.Mul(balance, c.MonthlyInterest)
	p := NewPayment(FP_PRECISION)
	p.PaymentNumber = len(c.PaymentsSchedule) + 1
	p.Interest.Mul(balance, c.MonthlyInterest)

	p.Principal.Sub(c.MonthlyPayment, interestPayment)
	p.Balance.Sub(balance, p.Principal)
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
	c := NewCustomerRecord(FP_PRECISION)
	c.Payments()
	PCSV(c.PaymentsSchedule)
}
