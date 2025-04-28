use rust_decimal::Decimal;
use rust_decimal::MathematicalOps;
use rust_decimal_macros::dec;
use std::fmt::Write as FmtWrite;

/// Represents a mortgage record from the input file
#[derive(Debug)]
struct MortgageRecord {
    id: String,
    surname: String,
    initial: char,
    first_name: String,
    loan_amount: Decimal,
    interest_rate: Decimal,
    term_years: u8,
    loan_type: char,
}

/// Represents a monthly payment
#[derive(Debug)]
struct PaymentRecord {
    year: u16,
    month: u8,
    payment_amount: Decimal,
    remaining_balance: Decimal,
}

/// A mortgage calculator that processes mortgage records
pub struct MortgageCalculator;

impl MortgageCalculator {
    /// Create a new mortgage calculator
    pub fn new() -> Self {
        MortgageCalculator
    }

    /// Process the input string buffer and return a formatted output string buffer
    pub fn process(&self, input_buffer: &str, start_year: u16) -> String {
        // Parse mortgage records from input string
        let records = self.parse_mortgage_records(input_buffer);

        // Process each record and build output string
        let mut output = String::new();

        for record in records {
            let payments = self.calculate_payment_schedule(&record, start_year);
            self.format_record_output(&record, &payments, &mut output);
        }

        output
    }

    /// Parse mortgage records from the input string
    fn parse_mortgage_records(&self, input: &str) -> Vec<MortgageRecord> {
        let mut records = Vec::new();

        for line in input.lines() {
            if line.len() < 48 {
                eprintln!("Warning: Skipping invalid line (too short): {}", line);
                continue;
            }

            let id = line[0..6].trim().to_string();
            let surname = line[6..20].trim().to_string();
            let initial = line.chars().nth(20).unwrap_or(' ');
            let first_name = line[21..35].trim().to_string();

            let loan_amount_str = line[35..41].trim();
            let loan_amount = match Decimal::from_str_exact(loan_amount_str) {
                Ok(amount) => amount,
                Err(e) => {
                    eprintln!("Warning: Invalid loan amount '{}': {}", loan_amount_str, e);
                    continue;
                }
            };

            let rate_str = format!("{}.{}", &line[41..43], &line[43..45]);
            let interest_rate = match Decimal::from_str_exact(&rate_str) {
                Ok(rate) => rate,
                Err(e) => {
                    eprintln!("Warning: Invalid interest rate '{}': {}", rate_str, e);
                    continue;
                }
            };

            let term_str = line[45..47].trim();
            let term_years = match term_str.parse::<u8>() {
                Ok(term) => term,
                Err(e) => {
                    eprintln!("Warning: Invalid term '{}': {}", term_str, e);
                    continue;
                }
            };

            let loan_type = line.chars().nth(47).unwrap_or(' ');

            records.push(MortgageRecord {
                id,
                surname,
                initial,
                first_name,
                loan_amount,
                interest_rate,
                term_years,
                loan_type,
            });
        }

        records
    }

    /// Calculate monthly payment schedule for a mortgage
    fn calculate_payment_schedule(&self, record: &MortgageRecord, start_year: u16) -> Vec<PaymentRecord> {
        // Early return if loan amount is zero
        if record.loan_amount.is_zero() {
            return Vec::new();
        }
        
        let mut payments = Vec::new();

        // Convert annual interest rate to monthly rate
        let monthly_rate = record.interest_rate / dec!(100) / dec!(12);

        // Number of payments
        let num_payments = record.term_years as u32 * 12;

        // Calculate monthly payment amount using the loan payment formula
        // P = L[c(1 + c)^n]/[(1 + c)^n - 1]
        // where P = payment, L = loan amount, c = monthly interest rate, n = number of payments
        let base = dec!(1) + monthly_rate;
        let power_term = base.powi(num_payments as i64);

        let numerator = monthly_rate * power_term;
        let denominator = power_term - dec!(1);

        // Handle potential division by zero if rate is 0 or term is 0
        let monthly_payment = if denominator.is_zero() {
            if record.loan_amount.is_zero() {
                dec!(0)
            } else if record.interest_rate.is_zero() && num_payments > 0 {
                record.loan_amount / Decimal::from(num_payments)
            } else {
                eprintln!("Warning: Cannot calculate payment for 0% interest and 0 term, or potential division by zero. Record ID: {}", record.id);
                dec!(0) 
            }
        } else {
            record.loan_amount * numerator / denominator
        };

        // Calculate payment schedule
        let mut remaining_balance = record.loan_amount;

        for payment_num in 1..=num_payments {
            let year = start_year + ((payment_num - 1) / 12) as u16;
            let month = ((payment_num - 1) % 12 + 1) as u8;

            // Calculate interest for this month
            let interest = remaining_balance * monthly_rate;

            // Calculate principal payment for this month
            let principal_payment = if monthly_payment <= remaining_balance {
                monthly_payment - interest
            } else {
                remaining_balance
            };

            // Calculate actual payment
            let actual_payment = interest + principal_payment;

            // Update remaining balance
            remaining_balance -= principal_payment;

            payments.push(PaymentRecord {
                year,
                month,
                payment_amount: actual_payment,
                remaining_balance,
            });

            // If balance is near zero, we're done
            if remaining_balance < dec!(0.01) {
                break;
            }
        }

        payments
    }

    /// Format a record and its payments into the output string
    fn format_record_output(&self, record: &MortgageRecord, payments: &[PaymentRecord], output: &mut String) {
        writeln!(
            output,
            "Mortgage ID: {}, Customer: {} {} {}",
            record.id, record.first_name, record.initial, record.surname
        ).unwrap();

        writeln!(
            output,
            "Loan Amount: ${}, Interest Rate: {}%, Term: {} years, Type: {}",
            record.loan_amount, record.interest_rate, record.term_years, record.loan_type
        ).unwrap();

        writeln!(output, "Payment Schedule:").unwrap();

        for payment in payments {
            writeln!(
                output,
                "Year: {}, Month: {}, Payment Amount: ${:.2}, Remaining Balance: ${}",
                payment.year, payment.month, payment.payment_amount, payment.remaining_balance
            ).unwrap();
        }

        writeln!(output).unwrap(); // Add a blank line between records
    }
}


#[cfg(test)]
mod tests {
    use super::*;
    use rust_decimal_macros::dec;

    #[test]
    fn test_parse_mortgage_records_valid() {
        let calculator = MortgageCalculator::new();
        let input = "000001Smith         JMichael       300000062515F";
        let records = calculator.parse_mortgage_records(input);

        assert_eq!(records.len(), 1);
        let record = &records[0];

        assert_eq!(record.id, "000001");
        assert_eq!(record.surname, "Smith");
        assert_eq!(record.initial, 'J');
        assert_eq!(record.first_name, "Michael");
        assert_eq!(record.loan_amount, dec!(300000));
        assert_eq!(record.interest_rate, dec!(6.25));
        assert_eq!(record.term_years, 15);
        assert_eq!(record.loan_type, 'F');
    }

    #[test]
    fn test_parse_mortgage_records_multiple() {
        let calculator = MortgageCalculator::new();
        let input = "000001Smith         JMichael       300000062515F\n\
                     000002Marqez        LAldo          280000045030V";
        let records = calculator.parse_mortgage_records(input);

        assert_eq!(records.len(), 2);

        // Check second record
        let record = &records[1];
        assert_eq!(record.id, "000002");
        assert_eq!(record.surname, "Marqez");
        assert_eq!(record.initial, 'L');
        assert_eq!(record.first_name, "Aldo");
        assert_eq!(record.loan_amount, dec!(280000));
        assert_eq!(record.interest_rate, dec!(4.50));
        assert_eq!(record.term_years, 30);
        assert_eq!(record.loan_type, 'V');
    }

    #[test]
    fn test_parse_mortgage_records_invalid_line_too_short() {
        let calculator = MortgageCalculator::new();
        let input = "TooShort";
        let records = calculator.parse_mortgage_records(input);

        // Invalid records should be skipped
        assert_eq!(records.len(), 0);
    }

    #[test]
    fn test_parse_mortgage_records_invalid_loan_amount() {
        let calculator = MortgageCalculator::new();
        let input = "000123SMITH         JJOHN           INVALIDR0025F";
        let records = calculator.parse_mortgage_records(input);

        // Invalid records should be skipped
        assert_eq!(records.len(), 0);
    }

    #[test]
    fn test_parse_mortgage_records_invalid_interest_rate() {
        let calculator = MortgageCalculator::new();
        let input = "000123SMITH         JJOHN           300000XXXX25F";
        let records = calculator.parse_mortgage_records(input);

        // Invalid records should be skipped
        assert_eq!(records.len(), 0);
    }

    #[test]
    fn test_parse_mortgage_records_invalid_term() {
        let calculator = MortgageCalculator::new();
        let input = "000123SMITH         JJOHN           3000000800XXF";
        let records = calculator.parse_mortgage_records(input);

        // Invalid records should be skipped
        assert_eq!(records.len(), 0);
    }

    #[test]
    fn test_calculate_payment_schedule_standard_mortgage() {
        let calculator = MortgageCalculator::new();

        let record = MortgageRecord {
            id: "000123".to_string(),
            surname: "SMITH".to_string(),
            initial: 'J',
            first_name: "JOHN".to_string(),
            loan_amount: dec!(100000),
            interest_rate: dec!(6.00),
            term_years: 30,
            loan_type: 'F',
        };
        
        let start_year = 2025;

        let payments = calculator.calculate_payment_schedule(&record, start_year);

        // There should be 30 years * 12 months = 360 payments
        assert_eq!(payments.len(), 360);

        // Check first payment
        let first = &payments[0];
        assert_eq!(first.year, start_year);
        assert_eq!(first.month, 1);

        // For a $100,000 loan at 6% over 30 years, the payment should be around $599.55
        let expected_payment = dec!(599.55);
        let difference = (first.payment_amount - expected_payment).abs();
        assert!(difference <= dec!(5.0), "Payment amount {} differs from expected {} by more than $5", first.payment_amount, expected_payment);

        // Verify last payment
        let last = &payments[payments.len() - 1];
        assert_eq!(last.year, start_year + 29);
        assert_eq!(last.month, 12);

        // Remaining balance should be very close to zero at the end
        assert!(last.remaining_balance <= dec!(0.01));
    }

    #[test]
    fn test_calculate_payment_schedule_zero_interest() {
        let calculator = MortgageCalculator::new();

        let record = MortgageRecord {
            id: "000123".to_string(),
            surname: "SMITH".to_string(),
            initial: 'J',
            first_name: "JOHN".to_string(),
            loan_amount: dec!(100000),
            interest_rate: dec!(0.00),
            term_years: 10,
            loan_type: 'F',
        };

        let start_year = 2025;

        let payments = calculator.calculate_payment_schedule(&record, start_year);

        // There should be 10 years * 12 months = 120 payments
        assert_eq!(payments.len(), 120);

        // For a $100,000 loan at 0% over 10 years, each payment should be exactly $833.33
        let expected_payment = dec!(833.33);
        let first = &payments[0];
        let difference = (first.payment_amount - expected_payment).abs();
        assert!(difference <= dec!(0.01), "Payment amount {} differs from expected {} by more than $0.01", first.payment_amount, expected_payment);

        // Balance should be zero after all payments
        let last = &payments[payments.len() - 1];
        assert!(last.remaining_balance <= dec!(0.01));
    }

    #[test]
    fn test_calculate_payment_schedule_zero_amount() {
        let calculator = MortgageCalculator::new();

        let record = MortgageRecord {
            id: "000123".to_string(),
            surname: "SMITH".to_string(),
            initial: 'J',
            first_name: "JOHN".to_string(),
            loan_amount: dec!(0),
            interest_rate: dec!(5.00),
            term_years: 10,
            loan_type: 'F',
        };

        let start_year = 2025;

        let payments = calculator.calculate_payment_schedule(&record, start_year);

        // For a $0 loan, we should have no payments
        assert_eq!(payments.len(), 0);
    }

    #[test]
    fn test_format_record_output() {
        let calculator = MortgageCalculator::new();

        let record = MortgageRecord {
            id: "000123".to_string(),
            surname: "SMITH".to_string(),
            initial: 'J',
            first_name: "JOHN".to_string(),
            loan_amount: dec!(100000),
            interest_rate: dec!(5.00),
            term_years: 10,
            loan_type: 'F',
        };

        let payments = vec![
            PaymentRecord {
                year: 2024,
                month: 1,
                payment_amount: dec!(1060.66),
                remaining_balance: dec!(98939.34),
            },
            PaymentRecord {
                year: 2024,
                month: 2,
                payment_amount: dec!(1060.66),
                remaining_balance: dec!(97873.46),
            },
        ];

        let mut output = String::new();
        calculator.format_record_output(&record, &payments, &mut output);

        // Check that the output contains key expected strings
        assert!(output.contains("Mortgage ID: 000123"));
        assert!(output.contains("Customer: JOHN J SMITH"));
        assert!(output.contains("Loan Amount: $100000"));
        assert!(output.contains("Interest Rate: 5.00%"));
        assert!(output.contains("Term: 10 years"));
        assert!(output.contains("Year: 2024, Month: 1"));
        assert!(output.contains("Payment Amount: $1060.66"));
        assert!(output.contains("Remaining Balance: $98939.34"));
    }

    #[test]
    fn test_process_end_to_end() {
        let calculator = MortgageCalculator::new();

        // Test with one valid record
        let input = "000001Smith         JMichael       300000062515F";
        let start_year = 2025;
        
        let output = calculator.process(input, start_year);

        // Check for expected output elements
        assert!(output.contains("Mortgage ID: 000001"));
        assert!(output.contains("Customer: Michael J Smith"));
        assert!(output.contains("Loan Amount: $300000"));
        assert!(output.contains("Year: 2025, Month: 1"));

        // There should be payment records for each month
        assert!(output.contains("Year: 2026, Month: 12"));
        assert!(output.contains("Year: 2027, Month: 1"));

        // The last payment should be near the end of the term
        assert!(output.contains("Year: 2033, Month: 12"));
    }

    #[test]
    fn test_process_with_invalid_records() {
        let calculator = MortgageCalculator::new();

        // Test with one valid and one invalid record
        let input = "000123SMITH         JJOHN           100000050010F\n\
                     INVALIDThis is not a valid record at all.";
        let start_year = 2025;
        
        let output = calculator.process(input, start_year);

        // Should only process the valid record
        assert!(output.contains("Mortgage ID: 000123"));
        assert!(!output.contains("INVALID"));
    }

    #[test]
    fn test_process_with_empty_input() {
        let calculator = MortgageCalculator::new();

        let input = "";
        let start_year = 2025;
        
        let output = calculator.process(input, start_year);

        // Should result in empty output
        assert!(output.is_empty());
    }
}
