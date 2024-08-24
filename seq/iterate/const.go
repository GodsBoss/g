package iterate

// Const creates an iterator that returns the same value forever.
func Const[First any](first First) func(yield func(First) bool) {
	return func(yield func(yieldFirst First) bool) {
		for yield(first) {
		}
	}
}
