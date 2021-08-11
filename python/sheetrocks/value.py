from enum import Enum

class ValueType(Enum):
    EMPTY = 0
    NUMBER = 1
    DATE = 2
    BOOLEAN = 3
    TEXT = 4
    ARRAY = 5

class Value:
    def __init__(valueType, number, text, date, boolean, array):
        self.valueType = valueType
        self.number = number
        self.text = text
        self.date = date
        self.boolean = boolean
        self.array = array