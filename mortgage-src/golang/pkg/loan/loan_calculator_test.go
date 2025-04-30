package loan_test

import (
	"testing"

	"mortgage/pkg/loan"

	"github.com/stretchr/testify/assert"
)

func TestMortgageCalculator(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate expected payment, interest and balance for one month", func(t *testing.T) {
		// Arrange
		principal := 2000
		annualPercentualRate := int(0.01555 * 12 * 100 * 100)
		oneMonth := 1
		loanCalculator := loan.NewCalculator(principal, annualPercentualRate, oneMonth)

		// Act
		result := loanCalculator.CalculateMonthlyPayment(20)

		// Assert
		assert.Equal(t, []loan.LoanPaymentSummary{
			{
				PaymentAmount: "2031.10000000000000000000",
				Balance:       "0.00000000000000000000",
			},
		}, result)
	})
}
