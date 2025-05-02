package mortgage

import (
	"mortgage/pkg/loan"
)

type Mortgage struct {
	Mortgagee *Mortgagee
	Loan      loan.Calculator
}

// id SERIAL PRIMARY KEY,
// mortgage_id VARCHAR(6) NOT NULL,
// customer VARCHAR(24) NOT NULL,
// loan_amount INTEGER NOT NULL,
// annual_interest_rate DECIMAL(2,28) NOT NULL,
// term_in_years INTEGER NOT NULL,
// type VARCHAR(1),
// year INTEGER,
// month INTEGER,
// payment_number INTEGER NOT NULL,
// monthly_interest_rate DECIMAL(10,28) NOT NULL,
// monthly_interest_amount DECIMAL(10,28) NOT NULL,
// monthly_payment DECIMAL(10,28) NOT NULL,
// monthly_principal DECIMAL(10,28) NOT NULL,
// remaining_balance DECIMAL(10,28) NOT NULL
type MortgagePaymentPlan struct {
	ID                    string
	Customer              string
	LoanAmount            int
	AnnualInterestRate    string
	TermInYears           int
	Type                  string
	Year                  int
	Month                 int
	PaymentNumber         int
	MonthlyInterestRate   string
	MonthlyInterestAmount string
	MonthlyPayment        string
	MonthlyPrincipal      string
	RemainingBalance      string
}

func NewMortgage(mortgagee *Mortgagee, loanCalculator loan.Calculator) *Mortgage {
	return &Mortgage{
		Mortgagee: mortgagee,
		Loan:      loanCalculator,
	}
}

func (mortgage *Mortgage) GetMonthlyPaymentPlan(precision uint) []*MortgagePaymentPlan {
	loanPaymentSummary := mortgage.Loan.CalculateMonthlyPayment(precision)
	return convertToMortgagePaymentPlan(mortgage, loanPaymentSummary)
}

func convertToMortgagePaymentPlan(mortgage *Mortgage, loanPaymentSummary []loan.LoanPaymentSummary) []*MortgagePaymentPlan {
	mortgagePaymentPlan := make([]*MortgagePaymentPlan, len(loanPaymentSummary))

	for installment, payment := range loanPaymentSummary {
		mortgagePaymentPlan[installment] = &MortgagePaymentPlan{
			ID:                 mortgage.Mortgagee.ID,
			Customer:           mortgage.Mortgagee.GetFullName(),
			LoanAmount:         mortgage.Loan.GetPrincipal(),
			AnnualInterestRate: mortgage.Loan.GetAnnualInterestRateInPercentage(),
			TermInYears:        mortgage.Loan.GetLoanTermInYears(),
			MonthlyPayment:     payment.PaymentAmount,
			PaymentNumber:      installment + 1,
			RemainingBalance:   payment.Balance,
		}
	}

	return mortgagePaymentPlan
}
