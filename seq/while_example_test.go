package seq_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq"
)

func ExampleWhile() {
	numbers := func(yield func(int) bool) {
		n := 1
		for yield(n) {
			n++
		}
	}

	lessThan := func(limit int) func(int) bool {
		return func(n int) bool {
			return n < limit
		}
	}

	for n := range seq.While(lessThan(5))(numbers) {
		fmt.Println(n)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
