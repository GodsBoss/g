package iterate

// Const creates an iterator that returns the same value repeatedly.
//
// Infinite iterator. Reusable.
func Const[First any](first First) func(yield func(First) bool) {
	return func(yield func(yieldFirst First) bool) {
		for yield(first) {
		}
	}
}
