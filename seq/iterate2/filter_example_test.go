package iterate2_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleFilter() {
	sumIsEven := func(first int, second int) bool {
		return (first+second)%2 == 0
	}

	values := []int{1, 1, 2, 2, 3, 3}

	for i, n := range iterate2.Filter(sumIsEven)(slices.All(values)) {
		fmt.Printf("index %d + value %d = %d\n", i, n, i+n)
	}

	// Output:
	// index 1 + value 1 = 2
	// index 2 + value 2 = 4
	// index 5 + value 3 = 8
}
