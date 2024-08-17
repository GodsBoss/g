package pred_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq"
	"github.com/GodsBoss/g/seq/pred"
)

func ExampleNot() {
	numbers := func(maximum int) func(yield func(int) bool) {
		current := 0

		return func(yield func(int) bool) {
			for current < maximum {
				yield(current)
				current++
			}
		}
	}

	isEven := func(n int) bool {
		return n%2 == 0
	}

	for n := range seq.Filter(numbers(10), pred.Not(isEven)) {
		fmt.Println(n)
	}

	// Output:
	// 1
	// 3
	// 5
	// 7
	// 9
}

func ExampleAllOf() {
	numbers := func(maximum int) func(yield func(int) bool) {
		current := 0

		return func(yield func(int) bool) {
			for current < maximum {
				yield(current)
				current++
			}
		}
	}

	isDivisibleBy := func(divisor int) func(int) bool {
		return func(n int) bool {
			return n%divisor == 0
		}
	}

	for n := range seq.Filter(numbers(13), pred.AllOf(isDivisibleBy(2), isDivisibleBy(3))) {
		fmt.Println(n)
	}

	// Output:
	// 0
	// 6
	// 12
}

func ExampleAnyOf() {
	numbers := func(maximum int) func(yield func(int) bool) {
		current := 0

		return func(yield func(int) bool) {
			for current < maximum {
				yield(current)
				current++
			}
		}
	}

	isDivisibleBy := func(divisor int) func(int) bool {
		return func(n int) bool {
			return n%divisor == 0
		}
	}

	for n := range seq.Filter(numbers(10), pred.AnyOf(isDivisibleBy(2), isDivisibleBy(3))) {
		fmt.Println(n)
	}

	// Output:
	// 0
	// 2
	// 3
	// 4
	// 6
	// 8
	// 9
}
