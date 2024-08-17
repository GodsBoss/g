package iterate2

import "iter"

// Concatenate takes multiple sequences and concatenates them.
//
// Note: If any of the sequences is an infinite sequence,
// every sequence after that will never have taken values from.
func Concatenate[First any, Second any](sequences ...iter.Seq2[First, Second]) iter.Seq2[First, Second] {
	return func(yield func(first First, second Second) bool) {
		cont := true
		for i := range sequences {
			if !cont {
				return
			}

			sequences[i](
				func(first First, second Second) bool {
					cont = yield(first, second)
					return cont
				},
			)
		}
	}
}
