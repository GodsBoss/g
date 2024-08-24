package iterate2

// Empty is an iterator which yields no values at all.
//
// Obviously finite and reusable.
func Empty[First any, Second any](_ func(First, Second) bool) {}
