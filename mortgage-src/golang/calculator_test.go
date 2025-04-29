package mortgage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T){

	t.Parallel()

	t.Run("Should calculate the monthly payment using float64", func(t *testing.T) {
		// Arrange
		principal := 10000
		annualPercentualRate := 350
		remainingMonths := 60

		// Act
		result := CalculateMonthlyPaymentFloat64(principal, annualPercentualRate, remainingMonths)

		// Assert
		assert.InDelta(t, 181.92, result, 0.01, "The monthly payment should be equal to the expected value")
	})
}
