package loan

import "math/big"

type LoanPaymentSummary struct {
	PaymentAmount string
	Balance       string
}

type Calculator interface {
	CalculateMonthlyPayment(precision uint) []LoanPaymentSummary
}

type Loan struct {
	Principal            int
	AnnualPercentualRate int
	RemainingMonths      int
}

func NewCalculator(principal int, annualPercentualRate int, remainingMonths int) Calculator {
	return &Loan{
		Principal:            principal,
		AnnualPercentualRate: annualPercentualRate,
		RemainingMonths:      remainingMonths,
	}
}

func (loan *Loan) CalculateMonthlyPayment(precision uint) []LoanPaymentSummary {
	mortgagePaymentSummary := []LoanPaymentSummary{}
	balance := new(big.Rat).SetInt64(int64(loan.Principal))

	for remainingMonths := loan.RemainingMonths; remainingMonths > 0; remainingMonths-- {
		payment, interest := CalculateMonthlyPaymentBigRat(balance, loan.AnnualPercentualRate, remainingMonths)
		balance = balance.Sub(balance.Add(balance, interest), payment)
		mortgagePaymentSummary = append(mortgagePaymentSummary, LoanPaymentSummary{
			PaymentAmount: payment.FloatString(int(precision)),
			Balance:       balance.FloatString(int(precision)),
		})
	}

	return mortgagePaymentSummary
}
