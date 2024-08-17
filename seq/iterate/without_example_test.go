package iterate_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleWithoutValues() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range iterate.WithoutValues(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// Peter
	// Paul
	// Mary
}

func ExampleWithoutKeys() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range iterate.WithoutKeys(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// 33
	// 22
	// 27
}
