package channel

import "iter"

func ToSequence[Value any](values <-chan Value) iter.Seq[Value] {
	return func(yield func(value Value) bool) {
		for valueFromChannel := range values {
			if !yield(valueFromChannel) {
				return
			}
		}
	}
}
