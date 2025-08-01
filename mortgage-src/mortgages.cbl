       IDENTIFICATION DIVISION.
       PROGRAM-ID. MORTGAGES.
       AUTHOR. AIKEDA.
       DATE-WRITTEN. 04/23/2025.
       DATE-COMPILED. 05/01/2025.

       ENVIRONMENT DIVISION.
       CONFIGURATION SECTION.
       SOURCE-COMPUTER. IBM-370.
       OBJECT-COMPUTER. IBM-370.
       INPUT-OUTPUT SECTION.
       FILE-CONTROL.
           SELECT MORTGAGE-FILE ASSIGN TO 'mortgagees.txt'
           ORGANIZATION IS LINE SEQUENTIAL
           FILE STATUS IS WS-FILE-STATUS.


       DATA DIVISION.
       FILE SECTION.
       FD  MORTGAGE-FILE
           LABEL RECORDS ARE STANDARD
           RECORD CONTAINS 48 CHARACTERS
           BLOCK CONTAINS 0 RECORDS.

       01  MORTGAGEE-RECORD.
           05  MORTGAGEE-ID              PIC 9(6).
           05  MORTGAGEE-NAME.
               10  MORTGAGEE-LAST-NAME   PIC X(14).
               10  MORTGAGEE-FIRST-INIT  PIC X.
               10  MORTGAGEE-FIRST-NAME  PIC X(14).
           05  MORTGAGEE-AMOUNT          PIC 9(6) VALUE 0.
           05  MORTGAGEE-RATE            PIC 9(2)V99 VALUE 0.
           05  MORTGAGEE-TERM            PIC 99 VALUE 0.
           05  MORTGAGEE-TYPE            PIC X.

       WORKING-STORAGE SECTION.
       01 CUST-DETAILS.
           05 CUST-NAME              PIC X(30).
       01 LOAN-VALUES.
           05 AMOUNT                 PIC 9(6) VALUE 0.
           05 RATE                   COMP-2 VALUE 0.
           05 YEARS                  PIC 9(2) VALUE 0.
           05 BALANCE                COMP-2 VALUE 0.

       01 REMAINING-MONTHS           PIC 9(3).

       01 PAYMENT-VALUES.
           05 PAYMENT-AMOUNT         COMP-2 VALUE 0.
           05 INTEREST-PAYMENT       COMP-2 VALUE 0.
           05 PRINCIPAL              COMP-2 VALUE 0.
           05 ACTUAL-AMOUNT          COMP-2 VALUE 0.
           05 RECALC-PAYMENT-AMOUNT  COMP-2 VALUE 0.
           05 OUT-ANN-RATE           PIC 9(1)V9(28) VALUE 0.
           05 OUT-PRINCIPAL          PIC ZZZZ.ZZZZZZZZZZZZZZZZZZZZZZZZZ.
           05 OUT-ACTUAL-AMOUNT      PIC ZZZZ.ZZZZZZZZZZZZZZZZZZZZZZZZZ.
           05 OUT-INTEREST-PAYMENT   PIC ZZZZ.ZZZZZZZZZZZZZZZZZZZZZZZZZ.
           05 OUT-PAYMENT-AMOUNT     PIC ZZZZ.ZZZZZZZZZZZZZZZZZZZZZZZZZ.
           05 OUT-BALANCE            PIC S9(6)V9(23) VALUE 0.
           05 ANN-RATE               COMP-2 VALUE 0.
           05 BASE                   COMP-2 VALUE 0.
           05 PAYMENT-NO             PIC 9(3)  VALUE 0.

       01  WS-FILE-STATUS            PIC XX.
       01  WS-SWITCHES.
           05  WS-EOF-SWITCH         PIC X VALUE 'N'.
               88  WS-EOF            VALUE 'Y'.
               88  WS-NOT-EOF        VALUE 'N'.

       01 COUNTERS.
           05 CURRENT-YEAR           PIC 9(4) VALUE 0.
           05 CURRENT-MONTH          PIC 9(2) VALUE 0.

       01  PRINT-LINE.
           05  LN-ID                 PIC X(6).
           05  FILLER                PIC X(16) VALUE ' Customer name: '.
           05  LN-CUST-NAME          PIC X(40).
           05  FILLER                PIC X(10) VALUE ' Amount: '.
           05  LN-LOAN-AMT           PIC ZZZZZZZ.
           05  FILLER                PIC X(10) VALUE ' Rate: '.
           05  LN-RATE               PIC ZZZZ.ZZZZZZZZZZ.

       PROCEDURE DIVISION.
       0000-MAIN.
           PERFORM 1000-INITIALIZE
           DISPLAY 'ID,Customer,Loan Amount,Interest Rate,Term,Type,'
                   'Year,Month,Payment No.,Monthly Rate,Interest,'
                   'Monthly Payment,Principal Payment,Actual Amount,'
                   'RemainingBalance'
           PERFORM 2000-PROCESS-FILE UNTIL WS-EOF
           PERFORM 3000-TERMINATE
           STOP RUN.
        
       1000-INITIALIZE.
           OPEN INPUT MORTGAGE-FILE
           
           IF WS-FILE-STATUS NOT = '00'
             DISPLAY 'ERROR OPENING MORTGAGE FILE: ' WS-FILE-STATUS
             STOP RUN
           END-IF
           PERFORM 1100-READ-FILE.
        
       1100-READ-FILE.
           READ MORTGAGE-FILE
               AT END SET WS-EOF TO TRUE
               NOT AT END SET WS-NOT-EOF TO TRUE
           END-READ.

       2000-PROCESS-FILE.

           PERFORM 2100-PROCESS-RECORD
           PERFORM 1100-READ-FILE.

       2100-PROCESS-RECORD.
           MOVE '                                        '
                    TO LN-CUST-NAME
           STRING MORTGAGEE-FIRST-NAME DELIMITED BY SPACE
                  ' '  DELIMITED BY SIZE
                  MORTGAGEE-FIRST-INIT DELIMITED BY SPACE
                  ' ' DELIMITED BY SIZE
                  MORTGAGEE-LAST-NAME DELIMITED BY SPACE
                  ' '  DELIMITED BY SIZE
             INTO LN-CUST-NAME
           END-STRING.

           MOVE MORTGAGEE-ID TO LN-ID
           MOVE MORTGAGEE-AMOUNT TO LN-LOAN-AMT
           MOVE MORTGAGEE-TERM TO YEARS
           MOVE MORTGAGEE-AMOUNT TO BALANCE
           MOVE MORTGAGEE-RATE TO RATE

           COMPUTE REMAINING-MONTHS = YEARS*12
           COMPUTE ANN-RATE ROUNDED = (RATE / 12) / 100
           COMPUTE BASE ROUNDED = (1+ANN-RATE) ** REMAINING-MONTHS

           COMPUTE PAYMENT-AMOUNT ROUNDED = MORTGAGEE-AMOUNT *
                     ( ( ANN-RATE * BASE  ) / (BASE - 1))

           ADD PAYMENT-AMOUNT TO ZERO GIVING OUT-PAYMENT-AMOUNT ROUNDED
      *    DISPLAY 'Mortgage ID: ' LN-ID ', Customer: '
      *         LN-CUST-NAME
      *     DISPLAY 'Loan Amount: $' LN-LOAN-AMT ', Interest Rate: '
      *         RATE '% Term: ' YEARS ' years'

           MOVE 2025 TO CURRENT-YEAR
           MOVE 1 TO CURRENT-MONTH
           MOVE 0 TO PAYMENT-NO

           PERFORM 2200-CALCULATE-MONTHLY-PAYMENT 
                         UNTIL BALANCE  < 0.


       2200-CALCULATE-MONTHLY-PAYMENT.
           ADD 1 TO PAYMENT-NO

           COMPUTE INTEREST-PAYMENT ROUNDED = ANN-RATE * BALANCE
           COMPUTE PRINCIPAL ROUNDED = PAYMENT-AMOUNT - INTEREST-PAYMENT
           COMPUTE ACTUAL-AMOUNT ROUNDED = PRINCIPAL + INTEREST-PAYMENT
           COMPUTE BALANCE = BALANCE - PRINCIPAL

           MOVE BALANCE TO OUT-BALANCE
           MOVE ACTUAL-AMOUNT TO OUT-ACTUAL-AMOUNT
           MOVE PRINCIPAL TO OUT-PRINCIPAL
           MOVE INTEREST-PAYMENT TO OUT-INTEREST-PAYMENT
           MOVE ANN-RATE TO OUT-ANN-RATE

           COMPUTE REMAINING-MONTHS = REMAINING-MONTHS - 1
           DISPLAY LN-ID ',' FUNCTION TRIM(LN-CUST-NAME) ','
                  MORTGAGEE-AMOUNT ','
                  RATE ',' MORTGAGEE-TERM ',' MORTGAGEE-TYPE ','
                  CURRENT-YEAR ',' CURRENT-MONTH ',' PAYMENT-NO ',' 
                  OUT-ANN-RATE ',' OUT-INTEREST-PAYMENT ',' 
                  OUT-PAYMENT-AMOUNT ',' OUT-PRINCIPAL ','
                  OUT-ACTUAL-AMOUNT ',' OUT-BALANCE

           

           IF CURRENT-MONTH = 12
               COMPUTE CURRENT-YEAR = CURRENT-YEAR + 1
               MOVE 1 TO CURRENT-MONTH
           ELSE
               COMPUTE CURRENT-MONTH = CURRENT-MONTH + 1
           END-IF.

       3000-TERMINATE.
           CLOSE MORTGAGE-FILE.

