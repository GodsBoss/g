package either

// NewFirst creates an either value wrapping a value of the first type.
func NewFirst[First any, Second any](first First) Value[First, Second] {
	return withItemFirst[First, Second]{
		item: first,
	}
}

type withItemFirst[First any, Second any] struct {
	item First
}

func (f withItemFirst[First, Second]) Invoke(ifFirst func(First), _ func(Second)) {
	ifFirst(f.item)
}

// GetFirst returns the wrapped first type value of an either value. It returns the zero value
// if the either value wrapped a value of the second type.
func GetFirst[First any, Second any](value Value[First, Second]) (item First, ok bool) {
	Invoke(
		value,
		func(first First) {
			item = first
			ok = true
		},
		nil,
	)

	return
}

// HasFirst checks whether the either value wraps a value of the first type.
func HasFirst[First any, Second any](value Value[First, Second]) bool {
	_, ok := GetFirst(value)
	return ok
}
