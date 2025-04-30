package mortgage

import (
	"fmt"
	"mortgage/internal/file"
	"mortgage/pkg/loan"
	"strconv"
	"strings"
)

type MortgageParser interface {
	Parse(filePath string) ([]*Mortgage, error)
}

type mortgageParse struct {
	linesReader *file.LinesReader
}

func (parser *mortgageParse) Parse(filePath string) ([]*Mortgage, error) {
	lines, err := parser.linesReader.GetLinesFrom(filePath)
	if err != nil {
		return nil, err
	}

	mortgages := []*Mortgage{}
	for _, line := range lines {
		fmt.Println(line)

		mortgagee := parseLineToMortgagee(line)

		loanCalculator, err := parseLineToLoan(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing loan: %w", err)
		}


		mortgages = append(mortgages, NewMortgage(mortgagee, loanCalculator))
	}

	return mortgages, nil
}

func parseLineToMortgagee(line string) *Mortgagee {
	id := line[0:6]
	lastName := strings.TrimSpace(line[6:20])
	middleInitial := strings.TrimSpace(line[20:21])
	firstName := strings.TrimSpace(line[21:35])
	return NewMortgagee(id, lastName, middleInitial, firstName)
}

func parseLineToLoan(line string) (loan.Calculator, error) {
	principalString := line[35:41]
	principal, err := strconv.ParseInt(principalString, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing principal: %w", err)
	}

	annualPercentualRateString := line[41:45]
	annualPercentualRate, err := strconv.ParseInt(annualPercentualRateString, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing annual percentual rate: %w", err)
	}

	loanTermInYearsString := line[45:47]
	loanTermInYears, err := strconv.ParseInt(loanTermInYearsString, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing loan term in years: %w", err)
	}
	remainingMonths := int(loanTermInYears) * 12

	return loan.NewCalculator(int(principal), int(annualPercentualRate), remainingMonths), nil
}

func NewParser() MortgageParser {
	return &mortgageParse{
		linesReader: file.NewLinesReader(),
	}
}

func NewParserFromFileSystem(fileSystem file.Opener) MortgageParser {
	return &mortgageParse{
		linesReader: file.NewLinesReaderFromFileSytem(fileSystem),
	}
}
