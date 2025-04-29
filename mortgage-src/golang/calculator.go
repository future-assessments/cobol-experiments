package mortgage

import (
	"math/big"
)

func CalculateMonthlyPaymentBigRat(principal int64, annualPercentualRate int64, remainingMonths int64) big.Rat {
	monthlyRate := big.NewRat(annualPercentualRate, 100*100*12)
	one := big.NewRat(1, 1)
	accumulationFactor := exponentialRat(new(big.Rat).Add(one, monthlyRate), remainingMonths)

	numerator := new(big.Rat).Mul(monthlyRate, accumulationFactor)
	denominator := new(big.Rat).Sub(accumulationFactor, one)

	principalRat := new(big.Rat).SetInt64(principal)
	monthlyPayment := new(big.Rat).Mul(principalRat, new(big.Rat).Quo(numerator, denominator))

	return *monthlyPayment
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
