package iterate2

import "iter"

// Map maps values from an iter.Seq into other values, given a conversion function.
func Map[FirstOutput any, SecondOutput any, Input any](convert func(Input) (FirstOutput, SecondOutput)) func(iter.Seq[Input]) iter.Seq2[FirstOutput, SecondOutput] {
	return func(sequence iter.Seq[Input]) iter.Seq2[FirstOutput, SecondOutput] {
		return func(yield func(FirstOutput, SecondOutput) bool) {
			sequence(
				func(input Input) bool {
					return yield(convert(input))
				},
			)
		}
	}
}

// Map maps values from an iter.Seq2 into other values, given a conversion function.
func Map2[FirstOutput any, SecondOutput any, FirstInput any, SecondInput any](convert func(FirstInput, SecondInput) (FirstOutput, SecondOutput)) func(iter.Seq2[FirstInput, SecondInput]) iter.Seq2[FirstOutput, SecondOutput] {
	return func(sequence iter.Seq2[FirstInput, SecondInput]) iter.Seq2[FirstOutput, SecondOutput] {
		return func(yield func(FirstOutput, SecondOutput) bool) {
			sequence(
				func(first FirstInput, second SecondInput) bool {
					return yield(convert(first, second))
				},
			)
		}
	}
}
