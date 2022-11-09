// Package either provides a value that wraps one of two values of different types. These types are
// called the "first" and the "second" type throughout the package.
//
// The package is written in a functional style, with top-level functions that revolve around
// a simple value type.
package either

// Value represents either an item of either type First or type Second.
type Value[First any, Second any] interface {
	// Invoke calls ifFirst if this either value wraps a First item,
	// if Second otherwise. Neither ifFirst nor ifSecond must be nil.
	// If no action is to be taken for either one (or both), use the package
	// Invoke function.
	Invoke(ifFirst func(First), ifSecond func(Second))
}

// Invoke is a helper function which safely calls value.Invoke with ifFirst and ifSecond guaranteed to be non-nil.
func Invoke[First any, Second any](value Value[First, Second], ifFirst func(First), ifSecond func(Second)) {
	if ifFirst == nil {
		ifFirst = func(First) {}
	}
	if ifSecond == nil {
		ifSecond = func(Second) {}
	}

	value.Invoke(ifFirst, ifSecond)
}

// Swap lets the types of an either value switch places. If the argument wrapped a value of the first type,
// the return value now wraps that value as the second type, and vice versa.
func Swap[First any, Second any](value Value[First, Second]) Value[Second, First] {
	var result Value[Second, First]

	value.Invoke(
		func(first First) {
			result = NewSecond[Second](first)
		},
		func(second Second) {
			result = NewFirst[Second, First](second)
		},
	)

	return result
}

// ToSlices returns all wrapped values as slices. If the argument wrapped a value of the first type,
// the first return value contains exactly that item, the second return value is empty.
// Otherwise, it's the other way around.
func ToSlices[First any, Second any](value Value[First, Second]) ([]First, []Second) {
	firsts := make([]First, 0)
	seconds := make([]Second, 0)

	value.Invoke(
		func(first First) {
			firsts = append(firsts, first)
		},
		func(second Second) {
			seconds = append(seconds, second)
		},
	)

	return firsts, seconds
}
