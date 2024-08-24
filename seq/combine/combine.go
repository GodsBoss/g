// Package combine takes a middleware-like approach for functions that take an iterator
// and return an iterator, a.k.a. transformers.
package combine

// Transformers combines a transformer from Input to Intermediate with a transformer from Intermediate to Output
// into a new transformer from Input to Output.
func Transformers[Input any, Intermediate any, Output any](inner func(Input) Intermediate, outer func(Intermediate) Output) func(Input) Output {
	return func(input Input) Output {
		return outer(inner(input))
	}
}
