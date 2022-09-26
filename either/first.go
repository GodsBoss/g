package either

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

func HasFirst[First any, Second any](value Value[First, Second]) bool {
	_, ok := GetFirst(value)
	return ok
}
