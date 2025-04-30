package calculator

import "math/big"

type MortgagePaymentSummary struct {
	PaymentAmount   string
	Balance   string
}

type MortgageCalculator interface {
	CalculateMonthlyPayment(precision uint) []MortgagePaymentSummary
}

type Mortgage struct {
	Principal            int
	AnnualPercentualRate int
	RemainingMonths      int
}

func NewMortgageCalculator(principal int, annualPercentualRate int, remainingMonths int) *Mortgage {
	return &Mortgage{
		Principal:            principal,
		AnnualPercentualRate: annualPercentualRate,
		RemainingMonths:      remainingMonths,
	}
}

func (mortgage *Mortgage) CalculateMonthlyPayment(precision uint) []MortgagePaymentSummary {

	mortgagePaymentSummary := []MortgagePaymentSummary{}
	balance := new(big.Rat).SetInt64(int64(mortgage.Principal))
	for remainingMonths := mortgage.RemainingMonths; remainingMonths > 0; remainingMonths-- {
		payment, interest := CalculateMonthlyPaymentBigRat(balance, mortgage.AnnualPercentualRate, remainingMonths)
		balance = balance.Sub(balance.Add(balance, interest), payment)
		mortgagePaymentSummary = append(mortgagePaymentSummary, MortgagePaymentSummary{
			PaymentAmount: payment.FloatString(int(precision)),
			Balance: balance.FloatString(int(precision)),
		})
	}


	return mortgagePaymentSummary
}
