package seq_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq"
)

func ExampleSwapKeyValues() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for age, name := range seq.SwapKeyValues(maps.All(ages)) {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Unordered output:
	// Peter is 33 years old.
	// Paul is 22 years old.
	// Mary is 27 years old.
}
