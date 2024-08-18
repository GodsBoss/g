package mathseq_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/mathseq"
)

func ExampleMax() {
	values := []int32{11, -17, 22, 5, -3, 18, 53, 35}

	maximum, ok := mathseq.Max(slices.Values(values))

	if ok {
		fmt.Printf("Maximum is %d.\n", maximum)
	}

	// Output:
	// Maximum is 53.
}

func ExampleMin() {
	values := []int32{11, -17, 22, 5, -3, 18, 53, 35}

	minimum, ok := mathseq.Min(slices.Values(values))

	if ok {
		fmt.Printf("Minimum is %d.\n", minimum)
	}

	// Output:
	// Minimum is -17.
}
