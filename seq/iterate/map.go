package iterate

import "iter"

// Map maps values from a sequence into other values, given a conversion function.
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
