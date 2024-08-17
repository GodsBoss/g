package iterate_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleConcatenate() {
	inputs := [][]string{
		{"foo", "bar", "baz"},
		{"abc", "xyz"},
		{"1", "2", "3"},
	}

	for s := range iterate.Concatenate(slices.Values(inputs[0]), slices.Values(inputs[1]), slices.Values(inputs[2])) {
		fmt.Println(s)
	}

	// Output:
	// foo
	// bar
	// baz
	// abc
	// xyz
	// 1
	// 2
	// 3
}
