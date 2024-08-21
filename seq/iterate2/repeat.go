package iterate2

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
func Repeat[First any, Second any](count int, reusable bool) func(iter.Seq2[First, Second]) iter.Seq2[First, Second] {
	// Without repetitions, we can just return an empty iterator.
	if count == 0 {
		return func(_ iter.Seq2[First, Second]) iter.Seq2[First, Second] {
			return func(_ func(_ First, _ Second) bool) {}
		}
	}

	// If the underlying sequence is not reusable, we iterate once while also storing all values in a slice. Iterators created
	// with slices.Values are reusable, so we can just call Repeat again, but this time with a reusable iterator and,
	// for finite repetitions, count is reduced by 1.
	if !reusable {
		return func(sequence iter.Seq2[First, Second]) iter.Seq2[First, Second] {
			return func(yield func(first First, second Second) bool) {
				running := true
				firstValues := make([]First, 0)
				secondValues := make([]Second, 0)
				sequence(
					func(first First, second Second) bool {
						firstValues = append(firstValues, first)
						secondValues = append(secondValues, second)
						running = yield(first, second)
						return running
					},
				)

				// We can't just call with count-1 here because count might be math.MinInt, which means infinite repetitions,
				// and count-1 would underflow to math.MaxInt, meaning finite repetitions.
				if count > 0 {
					count = count - 1
				}

				Repeat[First, Second](count, IsReusable)(Zip(slices.Values(firstValues), slices.Values(secondValues)))(yield)
			}
		}
	}

	return func(sequence iter.Seq2[First, Second]) iter.Seq2[First, Second] {
		return func(yield func(first First, second Second) bool) {
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
					func(first First, second Second) bool {
						iterationsInSequence++
						running = yield(first, second)
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

// Useful count constants to use with Repeat.
const (
	Never         = 0
	InfiniteTimes = -1
)

// Determine whether a sequence is reusable when calling Repeat.
const (
	IsReusable    = true
	IsNotReusable = false
)
