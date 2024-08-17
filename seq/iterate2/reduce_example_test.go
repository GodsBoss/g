package iterate2_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleReduce() {
	f := func(current int, s string, subtract bool) int {
		if subtract {
			return current - len(s)
		}

		return current + len(s)
	}

	inputs := map[string]bool{
		"foo":    false,
		"a":      true,
		"foobar": false,
		"xyz":    true,
	}

	fmt.Println(iterate2.Reduce(0, f)(maps.All(inputs)))

	// Output:
	// 5
}
