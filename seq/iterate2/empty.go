package iterate2

// Empty is an iterator which yields no values at all.
func Empty[First any, Second any](_ func(First, Second) bool) {}
