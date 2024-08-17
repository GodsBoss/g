package iterate_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleOverFirst() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range iterate.OverFirst(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// Peter
	// Paul
	// Mary
}

func ExampleOverSecond() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range iterate.OverSecond(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// 33
	// 22
	// 27
}
