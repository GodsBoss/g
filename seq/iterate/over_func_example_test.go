package iterate_test

import (
	"fmt"
	"math/rand"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleOverFunc() {
	var current int = 0
	counter := func() int {
		current++
		return current
	}

	for value := range iterate.OverFunc(counter) {
		if value > 3 {
			break
		}
		fmt.Println(value)
	}

	// Output:
	// 1
	// 2
	// 3
}

// ExampleOverFunc_rand prints random float64 values between 0 (inclusive) and 1 (exclusive) until
// a random value < 0.01 occurs.
func ExampleOverFunc_rand() {
	for n := range iterate.OverFunc(rand.New(rand.NewSource(666)).Float64) {
		if n < 0.01 {
			return
		}

		fmt.Println(n)
	}
}
