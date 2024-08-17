package iterate

import "iter"

// Filter yields values from a sequence for which a predicate holds true.
func Filter[Value any](predicate func(Value) bool) func(sequence iter.Seq[Value]) iter.Seq[Value] {
	return func(sequence iter.Seq[Value]) iter.Seq[Value] {
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
}
