package iterate

import "iter"

// ForEach invokes a function for every value from a sequence.
func ForEach[Value any](invoke func(Value)) func(iter.Seq[Value]) {
	return func(sequence iter.Seq[Value]) {
		for value := range sequence {
			invoke(value)
		}
	}
}
