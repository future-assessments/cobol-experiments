package mortgage

import (
	"mortgage/pkg/loan"
)

type Mortgage struct {
	Mortgagee *Mortgagee
	Loan      loan.Calculator
}

type MortgagePaymentPlan struct {
	ID             string
	MortgageeName  string
	MonthlyPayment string
	Balance        string
}

func NewMortgage(mortgagee *Mortgagee, loanCalculator loan.Calculator) *Mortgage {
	return &Mortgage{
		Mortgagee: mortgagee,
		Loan:      loanCalculator,
	}
}

func (mortgage *Mortgage) GetMonthlyPaymentPlan(precision uint) []MortgagePaymentPlan {
	loanPaymentSummary := mortgage.Loan.CalculateMonthlyPayment(precision)
	return convertToMortgagePaymentPlan(mortgage, loanPaymentSummary)
}

func convertToMortgagePaymentPlan(mortgage *Mortgage, loanPaymentSummary []loan.LoanPaymentSummary) []MortgagePaymentPlan {
	mortgagePaymentPlan := make([]MortgagePaymentPlan, len(loanPaymentSummary))

	for i, payment := range loanPaymentSummary {
		mortgagePaymentPlan[i] = MortgagePaymentPlan{
			ID:             mortgage.Mortgagee.ID,
			MortgageeName:  mortgage.Mortgagee.GetFullName(),
			MonthlyPayment: payment.PaymentAmount,
			Balance:        payment.Balance,
		}
	}

	return mortgagePaymentPlan
}
