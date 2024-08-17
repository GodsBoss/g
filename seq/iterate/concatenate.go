package iterate

import "iter"

// Concatenate takes multiple sequences and concatenates them.
//
// Note: If any of the sequences is an infinite sequence,
// every sequence after that will never have taken values from.
func Concatenate[Value any](sequences ...iter.Seq[Value]) iter.Seq[Value] {
	return func(yield func(value Value) bool) {
		cont := true
		for i := range sequences {
			if !cont {
				return
			}

			sequences[i](
				func(value Value) bool {
					cont = yield(value)
					return cont
				},
			)
		}
	}
}
