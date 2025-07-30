package com.example.mortgage;

import java.io.BufferedReader;

import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Objects;
import java.util.logging.ConsoleHandler;
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

            var line = fr.readLine();
            calculatePayments(line);

        } catch(IOException fnfe) {
            logger.severe(fnfe.getMessage());
        }
        return "";
    }

    private void calculatePayments(String customerData) {
        logger.info(String.format("Processing file: %s", customerData));
    }

    public static void main(String[] args) {
        MortgageCalculator calculator = new MortgageCalculator();
        calculator.processFile();
    }
}
