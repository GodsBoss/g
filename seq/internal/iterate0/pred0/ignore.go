package pred0

func IgnoreValue[Value any](predicate func() bool) func(_ Value) bool {
	return func(_ Value) bool {
		return predicate()
	}
}

func IgnoreValues[First any, Second any](predicate func() bool) func(_ First, _ Second) bool {
	return func(_ First, _ Second) bool {
		return predicate()
	}
}
