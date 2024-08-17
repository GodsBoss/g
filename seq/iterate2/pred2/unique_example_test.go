package pred2_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq/iterate2"
	"github.com/GodsBoss/g/seq/iterate2/pred2"
)

func ExampleUnique() {
	type input struct {
		first  string
		second string
	}

	inputs := []input{
		{
			first:  "foo",
			second: "bar",
		},
		{
			first:  "bar",
			second: "foo",
		},
		{
			first:  "foo",
			second: "bar",
		},
		{
			first:  "baz",
			second: "baz",
		},
	}

	iterator := func() func(yield func(first, second string) bool) {
		index := 0
		return func(yield func(first, second string) bool) {
			for index < len(inputs) {
				if !yield(inputs[index].first, inputs[index].second) {
					return
				}
				index++
			}
		}
	}

	for first, second := range iterate2.Filter(pred2.Unique[string, string]())(iterator()) {
		fmt.Printf("%s-%s\n", first, second)
	}

	// Output:
	// foo-bar
	// bar-foo
	// baz-baz
}
