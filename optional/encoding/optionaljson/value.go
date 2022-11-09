package optionaljson

import "encoding/json"

// Value represents an optional value that can marshaled to unmarshaled from JSON.
// The zero value is a valid empty value of T. Unlike the optional values created by
// NewWithItem() and NewEmpty(), this one can be altered, but only via Value.UnmarshalJSON().
type Value[T any] struct {
	t *T
}

// NewValue creates a new, non-empty JSON value. To create an empty value, just use Value's zero value.
func NewValue[T any](t T) Value[T] {
	return Value[T]{
		t: &t,
	}
}

// UnmarshalJSON marshals "null" into an empty optional value, every other string is marshaled into an item
// wrapped by this value.
func (value *Value[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		value.t = nil
		return nil
	}

	var t T
	err := json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	value.t = &t
	return nil
}

// MarshalJSON marshals the wrapped item if such item exists. Else, the marshaled JSON is "null".
func (value Value[T]) MarshalJSON() ([]byte, error) {
	if value.t == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*value.t)
}

// Invoke calls f with the item wrapped if this JSON value is not empty.
func (value Value[T]) Invoke(f func(item T)) {
	if value.t != nil {
		f(*value.t)
	}
}
