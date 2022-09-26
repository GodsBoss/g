package either_test

import (
	"fmt"

	"github.com/GodsBoss/g/either"
)

func ExampleGetFirst() {
	value := either.NewFirst[int, string](9001)

	first, ok := either.GetFirst(value)
	if ok {
		fmt.Printf("Stored: %d\n", first)
	}

	// Output:
	// Stored: 9001
}

func ExampleHasFirst() {
	value := either.NewFirst[int, string](666)

	if either.HasFirst(value) {
		fmt.Println("Has first value.")
	}

	// Output:
	// Has first value.
}
