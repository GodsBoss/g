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
