package iterate

// Empty is an iterator which yields no values at all.
//
// Obviously finite and reusable.
func Empty[Value any](_ func(Value) bool) {}
