package loan

import (
	"math/big"
)

func CalculateMonthlyPaymentBigRat(principal *big.Rat, annualPercentualRate *big.Rat, remainingMonths int) (*big.Rat, *big.Rat) {
	monthlyRate := new(big.Rat).Quo(annualPercentualRate, big.NewRat(12, 1))
	one := big.NewRat(1, 1)
	accumulationFactor := exponentialRat(new(big.Rat).Add(one, monthlyRate), int64(remainingMonths))

	numerator := new(big.Rat).Mul(monthlyRate, accumulationFactor)
	denominator := new(big.Rat).Sub(accumulationFactor, one)

	principalCopy := new(big.Rat).Set(principal)
	monthlyPayment := new(big.Rat).Mul(principalCopy, new(big.Rat).Quo(numerator, denominator))

	monthlyInterestAmount := principalCopy.Mul(principalCopy, monthlyRate)
	return monthlyPayment, monthlyInterestAmount
}

func exponentialRat(base *big.Rat, exponential int64) *big.Rat {
	one := big.NewRat(1, 1)
	if exponential == 0 {
		return one
	}

	baseCopy := new(big.Rat).Set(base)
	result := new(big.Rat).Set(one)
	for count := int64(0); count < exponential; count++ {
		result.Mul(result, baseCopy)
	}
	return result
}
