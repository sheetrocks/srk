package ext

import "srk"

func Name() string {
	return "ECHO"
}

func Help() string {
	return "# ECHO(arg1)\n\nEchos an input for testing purposes."
}

func Calculate(v []srk.Value) srk.Value {
	return v[0]
}
