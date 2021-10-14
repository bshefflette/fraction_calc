# Fraction Calculator

Handle simple fraction operations:

Problem:

Write a command-line program in the language of your choice that will take operations on fractions as an input and produce a fractional result.

Legal operators shall be \*, /, +, - (multiply, divide, add, subtract)

- Input/Output represented like:
  - Operands and operators shall be separated by one or more spaces like: `{Operand}{whitespace}{Operator}{whitespace}{Operand}`
  - Mixed number like: "3_1/4"
  - Whole Number like: "3"
  - Improper Fraction like: "9/8"
  - Negatives with leading "-"
- 2 operands 1 operator
- negatives are possible
- whole numbers are just '2'

Example run:

./fraction_calc "3 - 1"
-> 2

./fraction_calc "-3_2/9 - -1_44/8"
-> -9_13/18

./fraction*calc "-71*-7/-3 - -0\_-5/2"
-> -70_5/6
