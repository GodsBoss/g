// Package optional implements an optional value, also sometimes called "Maybe".
//
// The package is written in a functional style, with many top-level functions that
// revolve around a simple value type.
package optional

// Value represents an optional value. It may contain an item of type T.
type Value[T any] interface {
	// Invoke calls f only if an item is stored with that value.
	Invoke(f func(item T))
}

// HasItem checks whether the value actually contains an item. Opposite of IsEmpty().
func HasItem[T any](value Value[T]) bool {
	hasItem := false

	value.Invoke(
		func(_ T) {
			hasItem = true
		},
	)

	return hasItem
}

// IsEmpty checks whether the value is empty. Opposite of HasItem().
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
// The second return value signals whether a value could be extracted.
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

// NewWithItem creates an optional value that actually wraps an item. The returned value is immutable,
// i.e. cannot be changed to wrap a different value or no value at all.
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

// NewEmpty creates an empty optional value. The returned value is immutable, i.e. cannot be changed
// to contain a value.
func NewEmpty[T any]() Value[T] {
	return empty[T]{}
}

type empty[T any] struct{}

func (e empty[T]) Invoke(_ func(T)) {}

// Pointer wraps a pointer into an optional value. nil pointers become empty values,
// else a value containing an item is returned.
func Pointer[T any](t *T) Value[*T] {
	if t == nil {
		return NewEmpty[*T]()
	}
	return NewWithItem(t)
}

// Dereferences converts a pointer into an optional value of the referenced type.
// Nil pointers become empty values, else the pointer's dereferenced value is wrapped.
func Dereference[T any](t *T) Value[T] {
	if t == nil {
		return NewEmpty[T]()
	}
	return NewWithItem(*t)
}

// MapItem takes a map and a key, returning an empty value if that key does not exist
// in the map, and the corresponding item for that key otherwise.
// Only checks whether a map key exists and may happily wrap zero values (e.g. nil pointers).
func MapItem[Key comparable, Item any](m map[Key]Item, key Key) Value[Item] {
	item, ok := m[key]
	if ok {
		return NewWithItem(item)
	}
	return NewEmpty[Item]()
}

// SliceItem takes a slice and an index, returning an empty value if the index is out of
// bounds for the slice, and the corresponding item for that index otherwise.
// Only checks whether the index is in range and may happily wrap zero values (e.g. nil pointers).
func SliceItem[Item any](s []Item, index int) Value[Item] {
	if index >= 0 && index < len(s) {
		return NewWithItem(s[index])
	}
	return NewEmpty[Item]()
}

// FromOKResult wraps the result of a function that returns an Item and a boolean to indicate whether
// it returned a value. Returns an empty value if ok is false, else wraps item. Does not check for
// zero values, so a (nil, true) return value would result in a non-empty result.
func FromOKResult[Item any](item Item, ok bool) Value[Item] {
	if ok {
		return NewWithItem(item)
	}

	return NewEmpty[Item]()
}
