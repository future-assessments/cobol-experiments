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

### Overall Goal
Test out modern programming languages for floating point calculations and identify when rounding errors may cause issues when migrating away from COBOL

### Current Goal
1. Read an input file containing customer name, their loan amount, interest rate and loan term
2. Calculate the monthly payments and update the balance as the payment amount is identified
3. Print out each monthly payment until the Loan balance reaches $0

The file uses a column based format with each field starting at different columns in the file (see readme for the layout)

The file will be read as ASCII (not EBCDIC) and each row processed one by one printing out all months for each record.

**Assumptions**

- There is no need for a backend storage (file or database) to store progressive results, in-memory is fine.
- Speed is not a factor here, just accuracy
- The algorithm used to calculate payments will be the same across all languages

**Nice to Haves**

Show to difference in Loan Amount and Amount paid at the end of the monthly payments

**Extra Info**

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

