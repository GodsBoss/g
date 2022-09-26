package either_test

import (
	"fmt"

	"github.com/GodsBoss/g/either"
)

func ExampleGetSecond() {
	value := either.NewSecond[int, string]("Hello, world!")

	first, ok := either.GetSecond(value)
	if ok {
		fmt.Printf("Stored: %s\n", first)
	}

	// Output:
	// Stored: Hello, world!
}

func ExampleHasSecond() {
	value := either.NewSecond[int, string]("Nomen est omen.")

	if either.HasSecond(value) {
		fmt.Println("Has second value.")
	}

	// Output:
	// Has second value.
}
