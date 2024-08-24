package iterate

// Empty is an iterator which yields no values at all.
func Empty[Value any](_ func(Value) bool) {}
