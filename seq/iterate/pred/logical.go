package pred

// Not inverts the given predicate.
func Not[Value any](predicate func(Value) bool) func(Value) bool {
	return func(value Value) bool {
		return !predicate(value)
	}
}

// AllOf creates a predicate that returns false if at least one of the given predicates returns false, else true.
func AllOf[Value any](predicates ...func(Value) bool) func(Value) bool {
	return func(value Value) bool {
		for i := range predicates {
			if !predicates[i](value) {
				return false
			}
		}

		return true
	}
}

// AnyOf creates a predicate that returns true if at least one of the given predicates returns true, else false.
func AnyOf[Value any](predicates ...func(Value) bool) func(Value) bool {
	return func(value Value) bool {
		for i := range predicates {
			if predicates[i](value) {
				return true
			}
		}

		return false
	}
}
