package iterate

import (
	"iter"
	"slices"
)

// Repeat repeats sequences.
//
// count determines how often the sequence is to be repeated. A value of 0 creates a
// sequence with no values, a negative value repeats forever. Pass NoRepetition or InfiniteRepetitions
// for enhanced readability.
//
// reusable determines whether the given iterator can be re-used. If not (and count is not 0), all values
// from the underlying iterator are stored. Warning: Can lead to unbounded memory usage in case of an infinite iterator.
func Repeat[Value any](count int, reusable bool) func(iter.Seq[Value]) iter.Seq[Value] {
	// Without repetitions, we can just return an empty iterator.
	if count == 0 {
		return func(_ iter.Seq[Value]) iter.Seq[Value] {
			return func(_ func(_ Value) bool) {}
		}
	}

	// If the underlying sequence is not reusable, we iterate once while also storing all values in a slice. Iterators created
	// with slices.Values are reusable, so we can just call Repeat again, but this time with a reusable iterator and,
	// for finite repetitions, count is reduced by 1.
	if !reusable {
		return func(sequence iter.Seq[Value]) iter.Seq[Value] {
			return func(yield func(value Value) bool) {
				running := true
				values := make([]Value, 0)
				sequence(
					func(value Value) bool {
						values = append(values, value)
						running = yield(value)
						return running
					},
				)

				// We can't just call with count-1 here because count might be math.MinInt, which means infinite repetitions,
				// and count-1 would underflow to math.MaxInt, meaning finite repetitions.
				if count > 0 {
					count = count - 1
				}

				Repeat[Value](count, IsReusable)(slices.Values(values))(yield)
			}
		}
	}

	return func(sequence iter.Seq[Value]) iter.Seq[Value] {
		return func(yield func(value Value) bool) {
			// With a reusable iterator, we can just loop over it again and again, until yield returns false
			// or count is reached.
			running := true

			for running && count != 0 {
				if count > 0 {
					count--
				}

				// iterationsInSequence counts the iterations in the underlying iterator. If this is zero,
				// that iterator does not produce any items, meaning yield will never be called, resulting
				// in an infinite loop that can't be escaped via break or return in the client code's for loop.
				// Therefore, we return when this happens.
				iterationsInSequence := 0

				sequence(
					func(value Value) bool {
						iterationsInSequence++
						running = yield(value)
						return running
					},
				)

				if iterationsInSequence == 0 {
					return
				}
			}
		}
	}
}

const (
	Never         = 0
	InfiniteTimes = -1
)

const (
	IsReusable    = true
	IsNotReusable = false
)
