package seq

import "iter"

// While yields values from a sequence as long as the given condition holds true.
func While[Value any](sequence iter.Seq[Value], condition func(Value) bool) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		sequence(
			func(value Value) bool {
				if condition(value) {
					return yield(value)
				}

				return false
			},
		)
	}
}
