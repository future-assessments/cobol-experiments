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

#### Calculating Mortgage Payments

https://www.rocketloans.com/learn/financial-smarts/how-to-calculate-monthly-payment-on-a-loan#:~:text=Monthly%20Payment%20=%20(P%20%C3%97%20r,month%20repayment%20term%20(n).

**Amortizing Loans**

Unlike an interest-only loan, an amortizing loan payment goes toward both the interest and principal amount. That means you’ll be paying off the loan in equal monthly installments over the repayment term.

The formula for calculating the monthly payment on an amortizing personal loan is:

`Monthly Payment = P ((r (1+r)n) ∕ ((1+r)n−1))`

Let’s use the previous example, but this time, the personal loan you get is amortizing. The principal (P) is $10,000, the APR is 3.5% and you have a 60-month repayment term (n). With this formula, “r” stands for the annual rate, not the APR. You can use these steps to find the monthly payment:

* Divide your APR by 12 months to get your annual interest rate (r). Divide 0.035 by 12 to get 0.002917.
* Fill out the formula. You can now plug your loan information into the above equation. You should have $10,000((0.002917(1+0.002917)60) ∕ ((1+0.002917)60−1)).
* Solve the equations inside the first set of parentheses. You should end up with $10,000((0.002917 × 1.00291760) ∕ (1.00291760−1).
* Solve the exponentials. Calculate 1.00291760 to get 1.190967. The formula is now $10,000((0.002917 × 1.190967) ∕ (1.190967−1)).
* Solve the equations in the second set of parentheses. First, multiply 0.002917 by 1.190967 to get 0.003474. Then you can subtract 1 from 1.190967 to get 0.190967 for the other half of the equation. Your formula should look like $10,000(0.003474 ∕ 0.190967).
* Divide the numbers in the final set of parentheses. Take 0.003474 divided by 0.190967 to get 0.018192.
* Multiply the loan principal by the total. You will then multiply $10,000 by 0.018192 to get your monthly payment, $181.92.
* At this point, you can also use a loan calculator to make an amortization schedule for your loan. This extra step can help you visualize how your loan will be repaid over the length of the term. 

| Year    | Starting Balance | Interest | Principal  | Final Balance |
| ------- | ---------------- | -------- | ---------- | ------------- |
| 1       | $ 10,000.00      | $ 320.31 | $ 1,862.73 | $ 8,137.30    |
| 2       | $  8,137.30      | $ 254.05 | $ 1,928.99 | $ 6,208.35    |
| 3       | $  6,208.35      | $ 185.45 | $ 1,997.59 | $ 4,210.79    |
| 4       | $  4,210.79      | $ 114.40 | $ 2,068.64 | $ 2,142.18    |
| 5       | $  2,142.18      | $  40.84 | $ 2,142.20 | $     0.00    |
