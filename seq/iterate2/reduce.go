package iterate2

import "iter"

// Reduce reduces values from a sequence into a single value.
//
// Do not pass infinite iterators as this would lead to Reduce never returning.
func Reduce[Result any, First any, Second any](initial Result, reducer func(current Result, first First, second Second) Result) func(iter.Seq2[First, Second]) Result {
	return func(sequence iter.Seq2[First, Second]) Result {
		current := initial

		for first, second := range sequence {
			current = reducer(current, first, second)
		}

		return current
	}
}
