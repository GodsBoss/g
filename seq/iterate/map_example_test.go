package iterate_test

import (
	"fmt"
	"maps"
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

func ExampleMap2() {
	input := map[int]bool{
		1: true,
		2: false,
		3: false,
	}

	convert := func(n int, b bool) string {
		return fmt.Sprintf("%d => %t", n, b)
	}

	for s := range iterate.Map2(convert)(maps.All(input)) {
		fmt.Println(s)
	}

	// Unordered output:
	// 1 => true
	// 2 => false
	// 3 => false
}
