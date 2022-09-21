package optional

// Value represents an optional value. It may contain an item of type T.
type Value[T any] interface {
	// Invoke calls f only if an item is stored with that value.
	Invoke(f func(item T))
}

// HasItem checks whether the value actually contains an item.
func HasItem[T any](value Value[T]) bool {
	hasItem := false

	value.Invoke(
		func(_ T) {
			hasItem = true
		},
	)

	return hasItem
}

// IsEmpty checks whether the value is empty.
func IsEmpty[T any](value Value[T]) bool {
	return !HasItem(value)
}

// Len returns the number of items. This is 1 for a value containing an item, else 0.
func Len[T any](value Value[T]) int {
	if HasItem(value) {
		return 1
	}
	return 0
}

// GetItem extracts the item from value if such an item exists. If not, the zero value for T is returned.
// The second return value reports a value could be extracted.
func GetItem[T any](value Value[T]) (T, bool) {
	var item T
	ok := false

	value.Invoke(
		func(itemFromValue T) {
			item = itemFromValue
			ok = true
		},
	)

	return item, ok
}

// ToSlice returns the slice with the items wrapped by the value. Contains at most 1 item.
func ToSlice[T any](value Value[T]) []T {
	slice := make([]T, 0, 1)

	value.Invoke(
		func(item T) {
			slice = append(slice, item)
		},
	)

	return slice
}

// IfElse calls onValue for values wrapping an item, passing said item. If the value is
// empty, onEmpty is called instead. Both onValue and onEmpty may be nil.
func IfElse[T any](value Value[T], onValue func(T), onEmpty func()) {
	called := false

	value.Invoke(
		func(t T) {
			called = true
			if onValue != nil {
				onValue(t)
			}
		},
	)

	if !called && onEmpty != nil {
		onEmpty()
	}
}

// NewWithItem creates an optional value that actually wraps an item.
func NewWithItem[T any](t T) Value[T] {
	return withItem[T]{
		t: t,
	}
}

type withItem[T any] struct {
	t T
}

func (v withItem[T]) Invoke(f func(T)) {
	f(v.t)
}

// NewEmpty creates an empty optional value.
func NewEmpty[T any]() Value[T] {
	return empty[T]{}
}

type empty[T any] struct{}

func (e empty[T]) Invoke(_ func(T)) {}

// FromPointer wraps a pointer into an optional value. nil pointers become empty values,
// else a value containing an item is returned.
func FromPointer[T any](t *T) Value[*T] {
	if t == nil {
		return NewEmpty[*T]()
	}
	return NewWithItem(t)
}
