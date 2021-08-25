package values

import (
	"fmt"
	"time"
)

type ValueType uint8

const (
	EMPTY   ValueType = 0
	NUMBER  ValueType = 1
	DATE    ValueType = 2
	BOOLEAN ValueType = 3
	TEXT    ValueType = 4
	ARRAY   ValueType = 5
	ERROR   ValueType = 6
)

type Value struct {
	Type    ValueType
	Number  float64
	Text    string
	Date    time.Time
	Boolean bool
	Array   [][]Value
}

func (v *Value) String() string {
	switch v.Type {
	case EMPTY:
		return ""
	case NUMBER:
		return fmt.Sprintf("%g", v.Number)
	case DATE:
		return v.Date.Format(time.RFC3339)
	case BOOLEAN:
		if v.Boolean {
			return "TRUE"
		} else {
			return "FALSE"
		}
	case TEXT:
		return v.Text
	case ERROR:
		return v.Text
	default:
		// don't handle arrays for now
		return ""
	}
}
