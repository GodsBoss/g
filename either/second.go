package either

// NewSecond creates an either value wrapping a value of the second type.
func NewSecond[First any, Second any](second Second) Value[First, Second] {
	return withItemSecond[First, Second]{
		item: second,
	}
}

type withItemSecond[First any, Second any] struct {
	item Second
}

func (s withItemSecond[First, Second]) Invoke(_ func(First), ifSecond func(Second)) {
	ifSecond(s.item)
}

// GetSecond returns the wrapped second type value of an either value. It returns the zero value
// if the either value wrapped a value of the first type.
func GetSecond[First any, Second any](value Value[First, Second]) (item Second, ok bool) {
	Invoke(
		value,
		nil,
		func(second Second) {
			item = second
			ok = true
		},
	)

	return
}

// HasSecond checks whether the either value wraps a value of the second type.
func HasSecond[First any, Second any](value Value[First, Second]) bool {
	_, ok := GetSecond(value)
	return ok
}
