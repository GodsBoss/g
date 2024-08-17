package iterate2_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleWhile() {
	sumIsLessThan := func(maximum int) func(first int, second int) bool {
		return func(first int, second int) bool {
			return first+second < maximum
		}
	}

	values := []int{2, 3, 5, 7, 11, 13, 17, 19}

	for i, n := range iterate2.While(sumIsLessThan(10))(slices.All(values)) {
		fmt.Printf("index %d + value %d = %d\n", i, n, i+n)
	}

	// Output:
	// index 0 + value 2 = 2
	// index 1 + value 3 = 4
	// index 2 + value 5 = 7
}
