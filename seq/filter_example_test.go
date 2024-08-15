package seq_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq"
)

func ExampleFilter() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isOdd := func(n int) bool {
		return n%2 == 1
	}

	for n := range seq.Filter(seq.WithoutKeys(slices.All(numbers)), isOdd) {
		fmt.Println(n)
	}

	// Output:
	// 1
	// 3
	// 5
	// 7
	// 9
}
