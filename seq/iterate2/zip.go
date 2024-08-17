package iterate2

import "iter"

// Zip takes two single-value sequences and zips them into a two-value sequence. If any of the sequences ends, the resulting
// sequence ends as well. In the case that one sequence ends and the other doesn't, one value from the sequence that did not
// end vanishes.
func Zip[First any, Second any](firstSequence iter.Seq[First], secondSequence iter.Seq[Second]) iter.Seq2[First, Second] {
	pullFirst, cancelFirst := iter.Pull(firstSequence)
	pullSecond, cancelSecond := iter.Pull(secondSequence)

	return func(yield func(First, Second) bool) {
		defer cancelFirst()
		defer cancelSecond()

		for {
			first, firstOK := pullFirst()
			second, secondOK := pullSecond()

			if !firstOK || !secondOK {
				return
			}

			if !yield(first, second) {
				return
			}
		}
	}
}
