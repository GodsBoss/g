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
