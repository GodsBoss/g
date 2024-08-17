package iterate2

import "iter"

// Swap takes a iter.Seq2 and returns a new iter.Seq2 with first and second values swapped.
func Swap[First, Second any](sequence iter.Seq2[First, Second]) iter.Seq2[Second, First] {
	return func(yield func(second Second, first First) bool) {
		sequence(
			func(first First, second Second) bool {
				return yield(second, first)
			},
		)
	}
}
