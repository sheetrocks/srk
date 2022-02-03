from enum import Enum
from typing import Any
from sheetrocks import value as v

# Calculates the sum of all args and adds 1, including ranges.
# Ignores non-number values

# Ranges are always 2d lists. The lists are rows and the items are cells.

# The range `A1:C3` as seen in a sheet:
         
#          A      B       C
#    ------------------------
#    |
# 1  |    10      5       7
# 2  |     8      Bob     20
# 3  |    11      Sally   4

# The range `A1:C3` as a 2d list of lists passed to the `calculate` function:

# [
#  [10, 5, 7]
#  [8, Bob, 20]
#  [11, Sally, 4]
# ]


def calculate(*args) -> float:
    sum = 0.0
    for arg in args:
        if isinstance(arg, float):
            sum += arg
        if isinstance(arg, list):
            for row in arg:
                if isinstance(row, list):
                    for cell in row:
                        if isinstance(cell, float):
                            sum += cell
    return sum + 1