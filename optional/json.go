package optional

import "encoding/json"

// JSONValue represents an optional value that can marshaled to unmarshaled from JSON.
// The zero value is a valid empty value of T. Unlike the optional values created by
// NewWithItem() and NewEmpty(), this one can be altered, but only via JSONValue.UnmarshalJSON().
type JSONValue[T any] struct {
	t *T
}

// NewJSONValue creates a new, non-empty JSON value. To create an empty value, just use JSONValue's zero value.
func NewJSONValue[T any](t T) JSONValue[T] {
	return JSONValue[T]{
		t: &t,
	}
}

// UnmarshalJSON marshals "null" into an empty optional value, every other string is marshaled into an item
// wrapped by this value.
func (value *JSONValue[T]) UnmarshalJSON(data []byte) error {
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
func (value JSONValue[T]) MarshalJSON() ([]byte, error) {
	if value.t == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*value.t)
}

// Invoke calls f with the item wrapped if this JSON value is not empty.
func (value JSONValue[T]) Invoke(f func(item T)) {
	if value.t != nil {
		f(*value.t)
	}
}
