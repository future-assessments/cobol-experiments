package calculator

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {

	t.Parallel()

	t.Run("Should calculate the monthly payment using Big Rational", func(t *testing.T) {
		// Arrange
		principal := new(big.Rat).SetInt64(10000)
		annualPercentualRate := 350
		remainingMonths := 60

		// Act
		paymentAmount, interest := CalculateMonthlyPaymentBigRat(principal, annualPercentualRate, remainingMonths)

		// Assert
		assert.Equal(t, "181.92", paymentAmount.FloatString(2), "The monthly payment should be equal to the expected value")
		assert.Equal(t, "29.17", interest.FloatString(2), "The interest should be equal to the expected value")
	})
}
