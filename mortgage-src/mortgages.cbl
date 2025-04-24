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
           05  MORTGAGEE-AMOUNT          PIC 9(6).
           05  MORTGAGEE-RATE            PIC 9(2)V99.
           05  MORTGAGEE-TERM            PIC 99.
           05  MORTGAGEE-TYPE            PIC X.

       WORKING-STORAGE SECTION.
       01 CUST-DETAILS.
           05 CUST-NAME              PIC X(30).
       01 LOAN-VALUES.
           05 AMOUNT                 PIC 9(6).
           05 RATE                   COMP-2 VALUE 0.
           05 YEARS                  PIC 9(2).
           05 BALANCE                PIC 9(6)V99.

       01 PAYMENT-VALUES.
           05 PAYMENT-AMOUNT         PIC 9(5)V99.

       01  WS-FILE-STATUS            PIC XX.
       01 WS-SWITCHES.
           05  WS-EOF-SWITCH         PIC X VALUE 'N'.
               88  WS-EOF            VALUE 'Y'.
               88  WS-NOT-EOF        VALUE 'N'.

       01 PRINT-LINE.
           05  LN-ID                 PIC X(6).
           05  FILLER                PIC X(16) VALUE ' Customer name: '.
           05  LN-CUST-NAME          PIC X(40).
           05  FILLER                PIC X(10) VALUE ' Amount: '.
           05  LN-LOAN-AMT           PIC ZZZZZZZ.

       PROCEDURE DIVISION.
       0000-MAIN.
           PERFORM 1000-INITIALIZE
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
           MOVE MORTGAGEE-AMOUNT TO AMOUNT
           MOVE MORTGAGEE-RATE TO RATE

           DISPLAY 'ID        : ' MORTGAGEE-ID
           DISPLAY 'LAST NAME : ' MORTGAGEE-LAST-NAME
           DISPLAY 'INIT      : ' MORTGAGEE-FIRST-INIT
           DISPLAY 'FIRST NAME: ' MORTGAGEE-FIRST-NAME
           DISPLAY 'AMOUNT    : ' MORTGAGEE-AMOUNT
           DISPLAY 'RATE      : ' MORTGAGEE-RATE
           DISPLAY 'TERM      : ' MORTGAGEE-TERM
           DISPLAY 'TYPE      : ' MORTGAGEE-TYPE.

       2200-CALCULATE-MONTHLY-PAYMENT.
           COMPUTE PAYMENT-AMOUNT ROUNDED = (AMOUNT / (YEARS * 12) ) 
           COMPUTE BALANCE = AMOUNT - PAYMENT-AMOUNT

           DISPLAY 'NOT YET IMPLEMENTED'.

       3000-TERMINATE.
           CLOSE MORTGAGE-FILE.

