package either

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

func HasSecond[First any, Second any](value Value[First, Second]) bool {
	_, ok := GetSecond(value)
	return ok
}
