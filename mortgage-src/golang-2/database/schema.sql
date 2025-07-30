-- ID, MortgageID, Customer,Loan Amount,Interest Rate,Term,Type,Year,Month,Payment No.,Monthly Rate,Interest,Monthly Payment,Principal Payment,Actual Amount,RemainingBalance

CREATE TABLE IF NOT EXISTS mortgages (
  id SERIAL PRIMARY KEY,
  mortgage_id VARCHAR(6) NOT NULL,
  customer VARCHAR(24) NOT NULL,
  loan_amount INTEGER NOT NULL,
  annual_interest_rate DECIMAL(2,4) NOT NULL,
  term_in_years INTEGER NOT NULL,
  type VARCHAR(1),
  year INTEGER,
  month INTEGER,
  payment_number INTEGER NOT NULL,
  monthly_interest_rate DECIMAL(10,28) NOT NULL,
  monthly_interest_amount DECIMAL(10,28) NOT NULL,
  monthly_payment DECIMAL(10,28) NOT NULL,
  monthly_principal DECIMAL(10,28) NOT NULL,
  remaining_balance DECIMAL(10,28) NOT NULL
);
