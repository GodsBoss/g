package iterate

import "iter"

// Reduce reduces values from a sequence into a single value.
func Reduce[Value any](initial Value, reducer func(current Value, next Value) Value) func(iter.Seq[Value]) Value {
	return func(sequence iter.Seq[Value]) Value {
		current := initial

		for next := range sequence {
			current = reducer(current, next)
		}

		return current
	}
}
