// Package bitseq applies bitwise operations to sequences of integers.
package bitseq

import (
	"iter"
)

// Integer is a constraint for any of the builtin integer types.
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// And combines all items from sequence via bitwise AND. Returns -1 (all bits set to 1) if the sequence is empty.
// Stops iterating as soon as an intermediate value with all bits set to 0 is reached.
func And[T Integer](sequence iter.Seq[T]) T {
	var current T = T(minusOne)

	for next := range sequence {
		current = current & next
		if current == 0 {
			return current
		}
	}

	return current
}

// Or combines all items from sequence via bitwise OR. Returns 0 if the sequence is empty.
// Stops iterating as soon as an intermediate value with all bits set to 1 is reached.
func Or[T Integer](sequence iter.Seq[T]) T {
	var current T = 0

	for next := range sequence {
		current = current | next
		if current == T(minusOne) {
			return current
		}
	}

	return current
}

var minusOne int64 = -1
