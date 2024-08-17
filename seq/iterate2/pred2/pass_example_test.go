package pred2_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq/iterate2"
	"github.com/GodsBoss/g/seq/iterate2/pred2"
)

func ExamplePassFirstTo() {
	strLengthAtLeast := func(minimum int) func(s string) bool {
		return func(s string) bool {
			return len(s) >= minimum
		}
	}

	values := map[string]string{
		"baz":    "baz",
		"foo":    "foobar",
		"barfoo": "bar",
		"foobar": "barfoo",
	}

	for k, v := range iterate2.Filter(pred2.PassFirstTo[string, string](strLengthAtLeast(5)))(maps.All(values)) {
		fmt.Printf("%s => %s\n", k, v)
	}

	// Unordered output:
	// barfoo => bar
	// foobar => barfoo
}

func ExamplePassSecondTo() {
	strLengthAtLeast := func(minimum int) func(s string) bool {
		return func(s string) bool {
			return len(s) >= minimum
		}
	}

	values := map[string]string{
		"baz":    "baz",
		"foo":    "foobar",
		"barfoo": "bar",
		"foobar": "barfoo",
	}

	for k, v := range iterate2.Filter(pred2.PassSecondTo[string, string](strLengthAtLeast(5)))(maps.All(values)) {
		fmt.Printf("%s => %s\n", k, v)
	}

	// Unordered output:
	// foo => foobar
	// foobar => barfoo
}
