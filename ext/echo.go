package ext

import "github.com/sheetrocks/sheetrocks/srk/values"

func Name() string {
	return "ECHO"
}

func Help() string {
	return "# ECHO(arg1)\n\nEchos an input for testing purposes."
}

func Calculate(v []values.Value) values.Value {
	return v[0]
}
