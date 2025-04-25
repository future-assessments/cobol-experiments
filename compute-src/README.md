# Approach

4 Versions of the Calculator were created:
- COBOL
- Rust
- GoLang
- Java using float
- Java using BigDecimal

This is a simple calculation of the circumference of a circle (2 x &pi; x radius)

COBOL version was lifted from http://simotime.com/cblfpa01.htm

Pi is set as a value to a variable:
| Language        | Variable                                              |
| --------------- | ----------------------------------------------------- |
| Java BigDecimal | `var fpl_pi = new BigDecimal("3.14159265358979311");` |
| Java Float:     | `float fpl_pi = 3.14159265358979311f;`                |
| Rust:           | `let fpl_pi: f64 = 3.14159265358979323;`              |
| GoLang:         | `fpl_pi := 3.14159265358979323`                       |

<i>**Table 1.0** - Setting the value of &pi;</i>

# Outcomes

Results were pretty consistent except for the Java version using float as the data type:

| Language           | Circumference          | Pi                     | Radius             |
| ------------------ | ---------------------- | ---------------------- | ------------------ |
| COBOL              | 12.566370614359172     | 3.14159265358979311    | 2.000000000000000  |
| GoLang             | 12.566370614359172     | 3.141592653589793116   | 2.0000000000000000 |
| Rust               | 12.566370614359172     | 3.1415926535897931     | 2.0000000000000000 |
| Java (float)       | **12.566370964050293** | **3.141592741012573**  | 2.000000000000000  |
| Java (Big Decimal) | 12.566370614359172     | 3.141592653589793      | 2.000000000000000  |

<i>**Table 1.1** Output results</i>

You can see above that using Java floats, the rounding errors are apparent up front. Pi seems to be rounded at about the 7th decimal place with all following decimal places regenerated even though the value was specified from a string.

This leads to a value that is inconsistent with the other attempts.

So far for one off calculations all other languages produced the same result.

