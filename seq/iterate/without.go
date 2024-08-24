package iterate

import "iter"

// OverFirst creates a sequence of only the first values from a two-valued sequence.
//
// Finiteness and reusability depend on the underlying sequence.
func OverFirst[First, Second any](sequence iter.Seq2[First, Second]) iter.Seq[First] {
	return func(yield func(first First) bool) {
		sequence(
			func(first First, _ Second) bool {
				return yield(first)
			},
		)
	}
}

// OverSecond creates a sequence of only the second values from a two-valued sequence.
//
// Finiteness and reusability depend on the underlying sequence.
func OverSecond[First, Second any](sequence iter.Seq2[First, Second]) iter.Seq[Second] {
	return func(yield func(second Second) bool) {
		sequence(
			func(_ First, second Second) bool {
				return yield(second)
			},
		)
	}

}
