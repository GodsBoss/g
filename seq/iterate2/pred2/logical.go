package pred2

// Not inverts the given predicate.
func Not[First any, Second any](predicate func(First, Second) bool) func(First, Second) bool {
	return func(first First, second Second) bool {
		return !predicate(first, second)
	}
}

// AllOf creates a predicate that returns false if at least one of the given predicates returns false, else true.
func AllOf[First any, Second any](predicates ...func(First, Second) bool) func(First, Second) bool {
	return func(first First, second Second) bool {
		for i := range predicates {
			if !predicates[i](first, second) {
				return false
			}
		}

		return true
	}
}

// AnyOf creates a predicate that returns true if at least one of the given predicates returns true, else false.
func AnyOf[First any, Second any](predicates ...func(First, Second) bool) func(First, Second) bool {
	return func(first First, second Second) bool {
		for i := range predicates {
			if predicates[i](first, second) {
				return true
			}
		}

		return false
	}
}
