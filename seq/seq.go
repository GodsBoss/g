package seq

import "iter"

// WithoutValues takes a sequence of key/value pairs and returns a sequence of keys.
func WithoutValues[Key, Value any](sequence iter.Seq2[Key, Value]) iter.Seq[Key] {
	return func(yield func(key Key) bool) {
		sequence(
			func(key Key, _ Value) bool {
				return yield(key)
			},
		)
	}
}

// WithoutValues takes a sequence of key/value pairs and returns a sequence of values.
func WithoutKeys[Key, Value any](sequence iter.Seq2[Key, Value]) iter.Seq[Value] {
	return func(yield func(value Value) bool) {
		sequence(
			func(_ Key, value Value) bool {
				return yield(value)
			},
		)
	}

}

// SwapKeyValues takes a sequence of key/value pairs and returns a sequence of key/value pairs
// with keys and values being switched in their positions.
func SwapKeyValues[Key, Value any](sequence iter.Seq2[Key, Value]) iter.Seq2[Value, Key] {
	return func(yield func(value Value, key Key) bool) {
		sequence(
			func(key Key, value Value) bool {
				return yield(value, key)
			},
		)
	}
}
