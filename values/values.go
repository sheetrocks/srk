package values

import "time"

type ValueType uint8

const (
	EMPTY   ValueType = 0
	NUMBER  ValueType = 1
	DATE    ValueType = 2
	BOOLEAN ValueType = 3
	TEXT    ValueType = 4
	ARRAY   ValueType = 5
)

type Value struct {
	Type    ValueType
	Number  float64
	Text    string
	Date    time.Time
	Boolean bool
	Array   [][]Value
}
