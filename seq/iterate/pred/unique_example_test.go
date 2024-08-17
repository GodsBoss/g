package pred_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
	"github.com/GodsBoss/g/seq/iterate/pred"
)

func ExampleUnique() {
	input := []int{2, 7, 3, 3, 5, 2, 11, 7}

	for n := range iterate.Filter(pred.Unique[int]())(slices.Values(input)) {
		fmt.Println(n)
	}

	// Output:
	// 2
	// 7
	// 3
	// 5
	// 11
}
