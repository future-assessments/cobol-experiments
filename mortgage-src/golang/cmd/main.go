package main

import (
	"fmt"
	"log"
	"os"

	"mortgage/pkg/mortgage"

	"github.com/gocarina/gocsv"
)

const (
	precision = 20
)

func main() {
	mortgageFilePath, err := getFilePath()
	if err != nil {
		log.Fatalln("Error getting file path:", err)
		return
	}
	mortgageFileParser := mortgage.NewParser()

	mortgages, err := mortgageFileParser.Parse(mortgageFilePath)
	if err != nil {
		log.Fatalln("Error parsing file:", err)
		return
	}

	monthlyPaymentPlans := []*mortgage.MortgagePaymentPlan{}
	for _, mortgage := range mortgages {
		monthlyPaymentPlan := mortgage.GetMonthlyPaymentPlan(precision)
		for _, paymentPlan := range monthlyPaymentPlan {
			monthlyPaymentPlans = append(monthlyPaymentPlans, paymentPlan)
		}
	}

	csvContent, err := gocsv.MarshalString(&monthlyPaymentPlans)
		if err != nil {
			log.Fatalln("Error converting monthly payment plan to csv:", err)
			return
		}

		fmt.Println(csvContent)
}

func getFilePath() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	return "", fmt.Errorf("file path not provided")
}
