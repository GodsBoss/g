package seq

import "iter"

// While yields values from a sequence as long as the given predicate holds true.
func While[Value any](sequence iter.Seq[Value], predicate func(Value) bool) iter.Seq[Value] {
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
