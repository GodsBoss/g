package iterate2

// Const creates an iterator that returns the same value pair repeatedly.
//
// Infinite iterator.
func Const[First any, Second any](first First, second Second) func(yield func(First, Second) bool) {
	return func(yield func(yieldFirst First, yieldSecond Second) bool) {
		for yield(first, second) {
		}
	}
}
