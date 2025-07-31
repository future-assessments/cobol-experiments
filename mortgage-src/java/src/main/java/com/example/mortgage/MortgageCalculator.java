package com.example.mortgage;

import java.io.BufferedReader;

import java.io.IOException;
import java.io.InputStreamReader;
import java.math.BigDecimal;
import java.math.MathContext;
import java.math.RoundingMode;
import java.util.Objects;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;

public class MortgageCalculator {
    private final Logger logger;

    public MortgageCalculator() {
        logger = Logger.getLogger(MortgageCalculator.class.getName());
        logger.addHandler(new ConsoleHandler());
    }
    protected String processFile() {
        try (var fr = new BufferedReader(new InputStreamReader(
                Objects.requireNonNull(this.getClass().getResourceAsStream("/mortgages.txt"))))) {

            for (var line = fr.readLine(); line != null; line = fr.readLine()) {
                calculatePayments(line);
            }

        } catch(IOException fnfe) {
            logger.severe(fnfe.getMessage());
        }
        return "";
    }

    private void calculatePayments(String customerData) {
        logger.info(String.format("Processing file: %s", customerData));

        var clientId = customerData.substring(0, 6);
        var clientSurname = customerData.substring(6, 20).trim();
        var clientInitial = customerData.substring(20, 21);
        var clientFirstname = customerData.substring(21, 35).trim();
        String loanAmountString = customerData.substring(35, 41);
        logger.log(Level.INFO, "loanAmountString: {0}", new Object[]{loanAmountString});
        var loanInterestRateString = customerData.substring(41, 45);

        var termYears = Integer.parseInt(customerData.substring(45, 47));
        var clientName = String.format("%s %s %s", clientFirstname, clientInitial, clientSurname);
        logger.log(Level.INFO, "Client name {0}", new Object[]{clientName});

//        MathContext mc =  new MathContext(28, RoundingMode.CEILING);
        MathContext mc =  MathContext.DECIMAL128;
        var remainingMonths = termYears * 12;
        var interestRate = new BigDecimal(loanInterestRateString, mc).movePointLeft(2).setScale(28, RoundingMode.HALF_UP);
        var annualRate = (interestRate.divide(BigDecimal.valueOf(12), 28, RoundingMode.HALF_UP)).divide(BigDecimal.valueOf(100), 28, RoundingMode.HALF_UP);
        var base = annualRate.add(BigDecimal.valueOf(1), mc).pow(remainingMonths, mc).setScale(28, RoundingMode.HALF_UP);

        System.out.printf("loanTerm (months) %d%n", remainingMonths);
        System.out.printf("Interest rate: %2.8f%n", interestRate);
        System.out.printf("Interest rate/12: %6.24f%n",  interestRate.divide(BigDecimal.valueOf(12), mc));
        System.out.printf("annualRate: %6.24f%n",  annualRate);
        System.out.printf("base:  %6.24f%n", base);

        BigDecimal balance = new BigDecimal(loanAmountString, mc);
        System.out.printf("Balance %6.24f %n", balance);

        var amtPart1 = annualRate.multiply(base, mc);
        var amtPart2 = base.subtract(BigDecimal.valueOf(1), mc);
        var amtPart3 = amtPart1.divide(amtPart2, mc);

        BigDecimal paymentAmount = balance.multiply(amtPart3);
        System.out.printf("Payment amount: %6.24f %n", paymentAmount);
        var paymentNo = 1;

        System.out.printf("ID,Customer,Loan Amount,Interest Rate,Term,Type,Year,Month,Payment No.,Monthly Rate,Interest,Monthly Payment,Principal Payment,Actual Amount,RemainingBalance%n");;
        while (balance.compareTo(BigDecimal.ZERO) > 0) {
            var interestPayment = annualRate.multiply(balance, mc);
            var principal = paymentAmount.subtract(interestPayment, mc);
            var actualAmount = principal.add(interestPayment);

            balance = balance.subtract(actualAmount, mc);
            remainingMonths = remainingMonths - 1;
            // System.out.printf("Payment %d,%10.24f,%10.24f%n", paymentNo, principal, balance);
            System.out.printf("%s,%s,%s,%1.2f,%d,F,2025,00,%03d,%6.28f,%6.28f,%6.28f,%6.28f,%6.28f,%6.28f%n", clientId, clientName, loanAmountString, interestRate,
                    termYears, paymentNo, annualRate, interestPayment, paymentAmount, principal, actualAmount, balance);

            paymentNo++;
        }


    }

    public static void main(String[] args) {
        MortgageCalculator calculator = new MortgageCalculator();
        calculator.processFile();
    }
}
