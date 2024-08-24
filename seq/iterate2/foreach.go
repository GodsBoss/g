package iterate2

import "iter"

// ForEach invokes a function for every entry from a sequence.
//
// Returns only after exhausting the iterator. Does not return for infinite iterators.
func ForEach[First any, Second any](invoke func(First, Second)) func(iter.Seq2[First, Second]) {
	return func(sequence iter.Seq2[First, Second]) {
		for first, second := range sequence {
			invoke(first, second)
		}
	}
}
