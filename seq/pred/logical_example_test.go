package pred_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq"
	"github.com/GodsBoss/g/seq/pred"
)

func ExampleNot() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	isEven := func(n int) bool {
		return n%2 == 0
	}

	for n := range seq.Filter(slices.Values(numbers), pred.Not(isEven)) {
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
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	isDivisibleBy := func(divisor int) func(int) bool {
		return func(n int) bool {
			return n%divisor == 0
		}
	}

	for n := range seq.Filter(slices.Values(numbers), pred.AllOf(isDivisibleBy(2), isDivisibleBy(3))) {
		fmt.Println(n)
	}

	// Output:
	// 6
	// 12
}

func ExampleAnyOf() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	isDivisibleBy := func(divisor int) func(int) bool {
		return func(n int) bool {
			return n%divisor == 0
		}
	}

	for n := range seq.Filter(slices.Values(numbers), pred.AnyOf(isDivisibleBy(2), isDivisibleBy(3))) {
		fmt.Println(n)
	}

	// Output:
	// 2
	// 3
	// 4
	// 6
	// 8
	// 9
}
