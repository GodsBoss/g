package seq_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq"
)

func ExampleWithoutValues() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range seq.WithoutValues(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// Peter
	// Paul
	// Mary
}
