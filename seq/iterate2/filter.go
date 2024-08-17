package iterate2

import "iter"

// Filter yields values from a sequence for which a predicate holds true.
//
// See the pred2 subpackage for pre-defined predicates.
func Filter[First any, Second any](predicate func(First, Second) bool) func(iter.Seq2[First, Second]) iter.Seq2[First, Second] {
	return func(sequence iter.Seq2[First, Second]) iter.Seq2[First, Second] {
		return func(yield func(first First, second Second) bool) {
			sequence(
				func(first First, second Second) bool {
					if predicate(first, second) {
						return yield(first, second)
					}

					return true
				},
			)
		}
	}
}
