package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMortgageCalculator(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate expected payment, interest and balance for one month", func(t *testing.T) {
		// Arrange
		mortgage := NewMortgageCalculator(2000, 1866, 1)

		// Act
		result := mortgage.CalculateMonthlyPayment(20)

		// Assert
		assert.Equal(t, []MortgagePaymentSummary{
			{
				PaymentAmount: "2031.10000000000000000000",
				Balance: "0.00000000000000000000",
			},
		}, result)
	})
}
