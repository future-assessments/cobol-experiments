package mortgage

import (
	"math"
)

func CalculateMonthlyPaymentFloat64(principal int, annualPercentualRate int, remainingMonths int) float64 {
	monthlyRate := float64(annualPercentualRate) / 100.0 / 100.0 / 12.0
	accumulationFactor := math.Pow(1 + monthlyRate, float64(remainingMonths))

	numerator := monthlyRate * accumulationFactor
	denominator := accumulationFactor - 1

	monthlyPayment := float64(principal) * (numerator / denominator)

	return monthlyPayment
}
