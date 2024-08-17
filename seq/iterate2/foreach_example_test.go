package iterate2_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleForEach() {
	people := map[string]int{
		"Peter": 18,
		"Paul":  55,
		"Mary":  32,
	}

	output := func(name string, age int) {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	iterate2.ForEach(output)(maps.All(people))

	// Unordered output:
	// Peter is 18 years old.
	// Paul is 55 years old.
	// Mary is 32 years old.
}
