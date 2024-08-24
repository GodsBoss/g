package iterate

import "iter"

// Map maps values from an iter.Seq into other values, given a conversion function.
//
// Finiteness and reusability depend on the underlying sequence.
func Map[Output any, Input any](convert func(Input) Output) func(iter.Seq[Input]) iter.Seq[Output] {
	return func(sequence iter.Seq[Input]) iter.Seq[Output] {
		return func(yield func(Output) bool) {
			sequence(
				func(input Input) bool {
					return yield(convert(input))
				},
			)
		}
	}
}

// Map maps values from an iter.Seq2 into other values, given a conversion function.
//
// Finiteness and reusability depend on the underlying sequence.
func Map2[Output any, FirstInput any, SecondInput any](convert func(FirstInput, SecondInput) Output) func(iter.Seq2[FirstInput, SecondInput]) iter.Seq[Output] {
	return func(sequence iter.Seq2[FirstInput, SecondInput]) iter.Seq[Output] {
		return func(yield func(Output) bool) {
			sequence(
				func(first FirstInput, second SecondInput) bool {
					return yield(convert(first, second))
				},
			)
		}
	}
}
