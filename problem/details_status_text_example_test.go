package problem_test

import (
	"fmt"
	"net/http"

	"github.com/GodsBoss/g/problem"
)

func ExampleDetails_StatusText() {
	details := problem.Details{
		Status: http.StatusForbidden,
	}

	fmt.Println(details.StatusText())

	// Output:
	// Forbidden
}
