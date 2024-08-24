package iterate

import "iter"

// Reduce reduces values from a sequence into a single value.
//
// Do not pass infinite iterators as this would lead to Reduce never returning.
func Reduce[Result any, Value any](initial Result, reducer func(current Result, next Value) Result) func(iter.Seq[Value]) Result {
	return func(sequence iter.Seq[Value]) Result {
		current := initial

		for next := range sequence {
			current = reducer(current, next)
		}

		return current
	}
}
