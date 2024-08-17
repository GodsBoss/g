package iterate_test

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleMap() {
	input := []int{7, 666, 42}

	for s := range iterate.Map(strconv.Itoa)(slices.Values(input)) {
		fmt.Println(s)
	}

	// Output:
	// 7
	// 666
	// 42
}
