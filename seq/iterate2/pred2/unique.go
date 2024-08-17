package pred2

// Unique returns true for values that had not been passed before, false for others.
// This filter should be re-created instead of re-used as it holds a map of values that were passed to it.
// Memory usage grows with values passed to this. Must not be used by multiple sequences concurrently.
func Unique[First comparable, Second comparable]() func(First, Second) bool {
	type pair struct {
		first  First
		second Second
	}

	m := map[pair]struct{}{}

	return func(first First, second Second) bool {
		p := pair{
			first:  first,
			second: second,
		}

		_, ok := m[p]
		if ok {
			return false
		}

		m[p] = struct{}{}
		return true
	}
}
