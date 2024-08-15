package seq

import "iter"

// Filter yields values from a sequence for which a predicate holds true.
func Filter[Value any](sequence iter.Seq[Value], predicate func(Value) bool) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		sequence(
			func(value Value) bool {
				if predicate(value) {
					return yield(value)
				}

				return true
			},
		)
	}
}
