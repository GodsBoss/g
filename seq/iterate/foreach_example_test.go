package iterate_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleForEach() {
	names := []string{"Peter", "Paul", "Mary"}

	output := func(s string) {
		fmt.Println(s)
	}

	iterate.ForEach(output)(slices.Values(names))

	// Output:
	// Peter
	// Paul
	// Mary
}
