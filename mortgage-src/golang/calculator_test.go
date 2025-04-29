package mortgage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {

	t.Parallel()

	t.Run("Should calculate the monthly payment using Big Rational", func(t *testing.T) {
		// Arrange
		principal := int64(10000)
		annualPercentualRate := int64(350)
		remainingMonths := int64(60)

		// Act
		result := CalculateMonthlyPaymentBigRat(principal, annualPercentualRate, remainingMonths)

		// Assert
		assert.Equal(t, "181.92", result.FloatString(2), "The monthly payment should be equal to the expected value")
	})
}
