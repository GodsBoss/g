package iterate2

import "iter"

// While yields values from a sequence while a predicate holds true.
func While[First any, Second any](predicate func(First, Second) bool) func(iter.Seq2[First, Second]) iter.Seq2[First, Second] {
	return func(sequence iter.Seq2[First, Second]) iter.Seq2[First, Second] {
		return func(yield func(first First, second Second) bool) {
			sequence(
				func(first First, second Second) bool {
					if predicate(first, second) {
						return yield(first, second)
					}

					return false
				},
			)
		}
	}
}
