package problem_test

import (
	"fmt"

	"github.com/GodsBoss/g/problem"
)

func ExampleDetails_GetType_default() {
	details := problem.Details{}

	fmt.Printf("Problem details type is '%s'\n", details.GetType())

	// Output:
	// Problem details type is 'about:blank'
}

func ExampleDetails_GetType_populated() {
	details := problem.Details{
		Type: "https://example.com/problems/123",
	}

	fmt.Printf("Problem details type is '%s'\n", details.GetType())

	// Output:
	// Problem details type is 'https://example.com/problems/123'
}
