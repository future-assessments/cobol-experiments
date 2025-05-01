       IDENTIFICATION DIVISION.
       PROGRAM-ID. COMPUTEVALS.
       AUTHOR. AIKEDA.
       DATE-WRITTEN. 04/17/2025.
       DATE-COMPILED. 04/17/2025.

       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 X       PIC 9(3).
       01 xx      PIC 9(4).
       01 a       PIC 9(4)        VALUE 103.
       01 c       PIC 9(4)        VALUE 32.
       01 ZD-PI   PIC 9v9(17)     VALUE 0.
       01  FP-LONG.
           05  FPL-PI              COMP-2  value 0.
           05  FPL-RAD             COMP-2  value 0.
           05  FPL-CIR             COMP-2  value 0.
           05  FPL-SQUARE          PIC S9(3) USAGE IS COMP-3.
       01  PRINT-LINE.
           05  EDT-ID              pic X(3) value SPACES.
           05  FILLER              pic X(11) value ' Perimeter '.
           05  EDT-3-15-CIR        pic ZZZ.ZZZZZZZZZZZZZZZ.
           05  FILLER              pic X(08) value ' Radius '.
           05  EDT-3-15-RAD        pic ZZZ.ZZZZZZZZZZZZZZZ.
           05  FILLER              pic X(04) value ' Pi '.
           05  EDT-1-15-PI         pic Z.ZZZZZZZZZZZZZZZZZ.
           05  FILLER              pic X(05) value ' SQR '.
           05  EDT-1-15-SQR        pic Z.ZZZZZZZZZZZZZZZZZ.
       
       PROCEDURE DIVISION.
       MAIN.
           COMPUTE X = 24 * 3.
           DISPLAY X.

           DISPLAY "a is    : " a
           DISPLAY "c is    : " c

           COMPUTE xx = (a + 1) / c * 2.
           DISPLAY "xx is   : " xx

           MOVE 'FPL' to EDT-ID
           ADD 3.14159265358979323 TO ZERO GIVING FPL-PI ROUNDED
           ADD 2 TO ZERO GIVING FPL-RAD
           COMPUTE FPL-CIR ROUNDED = FPL-PI * (2 * FPL-RAD)
           COMPUTE FPL-SQUARE ROUNDED = FPL-RAD * FPL-RAD

           ADD FPL-CIR TO ZERO GIVING EDT-3-15-CIR
           ADD FPL-RAD TO ZERO GIVING EDT-3-15-RAD
           ADD FPL-PI  TO ZERO GIVING EDT-1-15-PI
           ADD FPL-SQUARE TO ZERO GIVING EDT-1-15-SQR

           DISPLAY PRINT-LINE UPON console
           STOP RUN.
