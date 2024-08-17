package pred

// Unique returns true for values that had not been passed before, false for others.
// This filter should be re-created instead of re-used as it holds a map of values that were passed to it.
// Memory usage grows with values passed to this. Must not be used by multiple sequences concurrently.
func Unique[Value comparable]() func(value Value) bool {
	m := map[Value]struct{}{}

	return func(value Value) bool {
		_, ok := m[value]
		if ok {
			return false
		}

		m[value] = struct{}{}
		return true
	}
}
