package formulas

import (
	"github.com/sheetrocks/srk/values"
)

func Echo(v []values.Value) values.Value {
	return v[0]
}
