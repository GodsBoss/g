package iterate_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleReduce() {
	add := func(n int, m int) int {
		return m + n
	}

	sum := iterate.Reduce(0, add)

	fmt.Println(sum(slices.Values([]int{2, 3, 5, 7})))

	// Output:
	// 17
}
