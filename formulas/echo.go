package formulas

import (
	"github.com/sheetrocks/srk/values"
)

func Echo(v []values.Value) values.Value {
	//return values.Value{Type: values.TEXT, Text: "SheetRocks is Amazing!"}
	return v[0]
}
