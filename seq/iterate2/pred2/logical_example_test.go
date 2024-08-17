package pred2_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate2"
	"github.com/GodsBoss/g/seq/iterate2/pred2"
)

func ExampleNot() {
	stringLengthMatches := func(str string, length int) bool {
		return len(str) == length
	}

	inputs := map[string]int{
		"foo":    3,
		"bar":    4,
		"baz":    5,
		"foobar": 6,
	}

	for str, length := range iterate2.Filter(pred2.Not(stringLengthMatches))(maps.All(inputs)) {
		fmt.Printf("'%s' does not have length %d.\n", str, length)
	}

	// Unordered output:
	// 'bar' does not have length 4.
	// 'baz' does not have length 5.
}

func ExampleAllOf() {
	sumIsEven := func(i, j int) bool {
		return (i+j)%2 == 0
	}

	firstIsGreater := func(i, j int) bool {
		return i > j
	}

	inputs := map[int]int{
		8: 4,
		4: 8,
		6: 3,
		3: 6,
	}

	for i, j := range iterate2.Filter(pred2.AllOf(sumIsEven, firstIsGreater))(maps.All(inputs)) {
		fmt.Printf("(%d, %d)\n", i, j)
	}

	// Unordered output:
	// (8, 4)
}

func ExampleAnyOf() {
	sumIsEven := func(i, j int) bool {
		return (i+j)%2 == 0
	}

	firstIsGreater := func(i, j int) bool {
		return i > j
	}

	inputs := map[int]int{
		4: 8,
		8: 4,
		6: 3,
		3: 6,
	}

	for i, j := range iterate2.Filter(pred2.AnyOf(sumIsEven, firstIsGreater))(maps.All(inputs)) {
		fmt.Printf("(%d, %d)\n", i, j)
	}

	// Unordered output:
	// (4, 8)
	// (8, 4)
	// (6, 3)
}
