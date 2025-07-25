package iterate2_test

import (
	"errors"
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleUntil() {
	results := slices.All(
		[]error{
			nil,
			nil,
			nil,
			errors.New("broken"),
			nil,
		},
	)

	errorOccured := func(_ int, err error) bool {
		return err != nil
	}

	for i := range iterate2.Until(errorOccured)(results) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
}
