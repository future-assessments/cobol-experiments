package main

import (
	"fmt"
	"log"

	"mortgage/pkg/mortgage"

	"github.com/gocarina/gocsv"
)

const (
	precision = 20
)

func main() {
	mortgageFilePath := "./tests/utils/examples/one-example.txt"
	mortgageFileParser := mortgage.NewParser()

	mortgages, err := mortgageFileParser.Parse(mortgageFilePath)
	if err != nil {
		log.Fatalln("Error parsing file:", err)
		return
	}

	for _, mortgage := range mortgages {
		monthlyPaymentPlan := mortgage.GetMonthlyPaymentPlan(precision)
		csvContent, err := gocsv.MarshalString(&monthlyPaymentPlan)
		if err != nil {
			log.Fatalln("Error converting monthly payment plan to csv:", err)
			return
		}

		fmt.Println(csvContent)
	}
}
