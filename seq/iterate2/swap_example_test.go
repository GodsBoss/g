package iterate2_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleSwapKeyValues() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for age, name := range iterate2.SwapKeyValues(maps.All(ages)) {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Unordered output:
	// Peter is 33 years old.
	// Paul is 22 years old.
	// Mary is 27 years old.
}
