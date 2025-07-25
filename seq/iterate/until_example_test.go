package iterate_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleUntil() {
	numbers := func(yield func(int) bool) {
		n := 1
		for yield(n) {
			n++
		}
	}

	greaterThan := func(limit int) func(int) bool {
		return func(n int) bool {
			return n > limit
		}
	}

	for n := range iterate.Until(greaterThan(3))(numbers) {
		fmt.Println(n)
	}

	// Output:
	// 1
	// 2
	// 3
}
