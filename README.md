# COBOL Validations

Code bases to test out different COBOL Constructs in different languages

## Compute

Basic floating point calculations in:
* COBOL
* RustLang
* GoLang
* Java using float
* Java using BigDecimal

Simple circumference calculation of &pi; * 2 * <i>radius</i>


## Mortgages

This program reads a file of customers with the following format

| Field | Size | Type |
| ------| ---- | ---- |
| ID | 6 characters | Integer (000001)|
| Surname | 14 characters | String |
| Initial | 1 character | String |
| First name | 14 characters | String |
| Loan Amount | 6 characters | Long (300000)|
| Loan rate | 4 characters | Decimal (00.00) |
| Loan Term | 2 characters | Integer (10) |
| Loan Type | 1 character | char |


The application will read the record and calculate the monthly payments and remaining balance printing out each monthly transaction:

```
Year: 2024, Month: 1, Payment Amount: $3456, Remaining Balance: $ $296544
Year: 2024, Month: 2, Payment Amount: $3456, Remaining Balance: $ $293088
```

We need to calculate the balance for the month based on the interest rate
We then need to calculate the monthly payment
We then subtract the paid amount from the overall balance

This is a simple calculation - don't include insurance and equity

