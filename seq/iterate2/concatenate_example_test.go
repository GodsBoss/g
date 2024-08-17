package iterate2_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleConcatenate() {
	inputs := [][]string{
		{"Peter", "Paul"},
		{"Mary", "Joannah"},
	}

	for i, s := range iterate2.Concatenate(slices.All(inputs[0]), slices.All(inputs[1])) {
		fmt.Printf("%d: %s\n", i, s)
	}

	// Output:
	// 0: Peter
	// 1: Paul
	// 0: Mary
	// 1: Joannah
}
