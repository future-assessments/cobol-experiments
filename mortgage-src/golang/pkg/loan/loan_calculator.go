package loan

import "math/big"

type LoanPaymentSummary struct {
	PaymentAmount string
	Balance       string
}

type Calculator interface {
	CalculateMonthlyPayment(precision uint) []LoanPaymentSummary
	GetPrincipal() int
	GetAnnualInterestRateInPercentage() string
	GetLoanTermInYears() int
}

type Loan struct {
	principal          int
	annualInterestRate *big.Rat
	remainingMonths    int
}

func NewCalculatorFromMonths(principal int, annualPercentualRate *big.Rat, remainingMonths int) Calculator {
	return &Loan{
		principal:          principal,
		annualInterestRate: annualPercentualRate,
		remainingMonths:    remainingMonths,
	}
}

func NewCalculator(principal int, annualPercentualRate *big.Rat, termInYears int) Calculator {
	return &Loan{
		principal:          principal,
		annualInterestRate: annualPercentualRate,
		remainingMonths:        termInYears * 12,
	}
}

func (loan *Loan) CalculateMonthlyPayment(precision uint) []LoanPaymentSummary {
	mortgagePaymentSummary := []LoanPaymentSummary{}
	balance := new(big.Rat).SetInt64(int64(loan.principal))

	for remainingMonths := loan.remainingMonths; remainingMonths > 0; remainingMonths-- {
		payment, interest := CalculateMonthlyPaymentBigRat(balance, loan.annualInterestRate, remainingMonths)
		balance = balance.Sub(balance.Add(balance, interest), payment)
		mortgagePaymentSummary = append(mortgagePaymentSummary, LoanPaymentSummary{
			PaymentAmount: payment.FloatString(int(precision)),
			Balance:       balance.FloatString(int(precision)),
		})
	}

	return mortgagePaymentSummary
}

func (loan *Loan) GetPrincipal() int {
	return loan.principal
}
func (loan *Loan) GetAnnualInterestRateInPercentage() string {
	return new(big.Rat).Mul(loan.annualInterestRate, big.NewRat(100, 1)).FloatString(2)
}

func (loan *Loan) GetLoanTermInYears() int {
	return loan.remainingMonths / 12
}
