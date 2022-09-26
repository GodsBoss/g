package optional

import "encoding/json"

// JSONValue represents an optional value that can marshaled to unmarshaled from JSON.
// The zero value is a valid empty value of T.
type JSONValue[T any] struct {
	t *T
}

// NewJSONValue creates a new, non-empty JSON value. To create an empty value, just use JSONValue's zero value.
func NewJSONValue[T any](t T) JSONValue[T] {
	return JSONValue[T]{
		t: &t,
	}
}

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

func (value JSONValue[T]) MarshalJSON() ([]byte, error) {
	if value.t == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*value.t)
}

func (value JSONValue[T]) Invoke(f func(item T)) {
	if value.t != nil {
		f(*value.t)
	}
}
