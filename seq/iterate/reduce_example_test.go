package iterate_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleReduce() {
	addStringLength := func(current int, next string) int {
		return current + len(next)
	}

	stringLengthSum := iterate.Reduce(0, addStringLength)

	fmt.Println(stringLengthSum(slices.Values([]string{"foo", "bar", "foobar"})))

	// Output:
	// 12
}
