# SUMPLUSONE

Returns 1 plus sum of all arguments given. If an argument is not a number, it ignores the value. Also accepts ranges such as A1:C3.

---
## Syntax

SUMPLUSONE(num, [nums])

#### num
A number to add to the total sum value.

#### [nums]
[ OPTIONAL ] - a range of numbers to add to the total sum value.


---
## Examples

EXAMPLE_RANGE
1, 2, 3

SUMPLUSONE(A1:C1)=7

SUMPLUSONE(1)=2

SUMPLUSONE(A1, B1, 5)=9

---
## Notes

- This formula is equivalent to using the standard SUM formula and adding 1. It is given for demonstrastion purposes and may not have practical value.

---
## Related Formulas

- [SUM](/formulas/sum): The SUM formula adds numbers. If an argument is a range, it will add all numbers in the range. If an argument or item in a range is not a number, it will be ignored.