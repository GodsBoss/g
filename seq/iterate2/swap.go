package iterate2

import "iter"

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
