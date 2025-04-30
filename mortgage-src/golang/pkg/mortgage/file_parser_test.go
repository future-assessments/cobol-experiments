package mortgage_test

import (
	"mortgage/pkg/loan"
	"mortgage/pkg/mortgage"
	utils_test "mortgage/tests/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMortgageParser(t *testing.T) {
	t.Parallel()

	t.Run("should parse a file and return a list of mortgages", func(t *testing.T) {
		// Arrange
		filePath := "examples/one-example.txt"
		parser := mortgage.NewParserFromFileSystem(utils_test.ExampleFilesFS)

		// Act
		mortgages, err := parser.Parse(filePath)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, []*mortgage.Mortgage{
			{
				Mortgagee: &mortgage.Mortgagee{
					ID:            "000002",
					LastName:      "Marqez",
					MiddleInitial: "L",
					FirstName:     "Aldo",
				},
				Loan: &loan.Loan{
					Principal:            280000,
					AnnualPercentualRate: 450,
					RemainingMonths:      30 * 12,
				},
			},
		}, mortgages)
	})
}
