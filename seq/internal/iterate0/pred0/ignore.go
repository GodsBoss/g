package pred0

// IgnoreValue takes a nonary predicate and converts it into a unary predicate which ignores its argument.
func IgnoreValue[Value any](predicate func() bool) func(_ Value) bool {
	return func(_ Value) bool {
		return predicate()
	}
}

// IgnoreValue takes a nonary predicate and converts it into a binary predicate which ignores its arguments.
func IgnoreValues[First any, Second any](predicate func() bool) func(_ First, _ Second) bool {
	return func(_ First, _ Second) bool {
		return predicate()
	}
}
