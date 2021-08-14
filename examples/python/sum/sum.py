from enum import Enum
from typing import Any
from sheetrocks import value as v

# Calculates the sum of all args, including ranges (which are always two dimensional lists)
# Ignores non-number values
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
    return sum