package iterate

import "iter"

// While yields values from a sequence as long as the given predicate holds true.
//
// See the pred subpackage for pre-defined predicates.
//
// Creates finite iterators from infinite iterators.
//
// Reusable if underlying sequence is reusable.
func While[Value any](predicate func(Value) bool) func(sequence iter.Seq[Value]) iter.Seq[Value] {
	return func(sequence iter.Seq[Value]) iter.Seq[Value] {
		return func(yield func(Value) bool) {
			sequence(
				func(value Value) bool {
					if predicate(value) {
						return yield(value)
					}

					return false
				},
			)
		}
	}
}
