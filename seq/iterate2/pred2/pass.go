package pred2

// PassFirstTo takes a one-valued predicate and converts it into a two-valued predicate, passing the first argument.
func PassFirstTo[First any, Second any](predicate func(First) bool) func(first First, _ Second) bool {
	return func(first First, _ Second) bool {
		return predicate(first)
	}
}

// PassSecondTo takes a one-valued predicate and converts it into a two-valued predicate, passing the second argument.
func PassSecondTo[First any, Second any](predicate func(Second) bool) func(_ First, second Second) bool {
	return func(_ First, second Second) bool {
		return predicate(second)
	}
}
