package iterate2_test

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleMap() {
	input := []string{"a", "baa", "aaaxxx"}

	// convert cuts of "a"s from the front of s, returns whether any were found and the length of the remaining string.
	convert := func(s string) (bool, int) {
		ok := strings.HasPrefix(s, "a")
		s = strings.TrimLeft(s, "a")
		return ok, len(s)
	}

	for ok, n := range iterate2.Map(convert)(slices.Values(input)) {
		if ok {
			fmt.Printf("Some 'a's were removed, rest has length %d.\n", n)
		} else {
			fmt.Printf("No 'a's were removed, length is %d.\n", n)
		}
	}

	// Output:
	// Some 'a's were removed, rest has length 0.
	// No 'a's were removed, length is 3.
	// Some 'a's were removed, rest has length 3.
}

func ExampleMap2() {
	input := map[string]int{
		"ab":  2,
		"xyz": 1,
		"foo": 3,
	}

	convert := func(s string, n int) (int, string) {
		s = strings.Repeat(s, n)
		return len(s), s
	}

	for n, s := range iterate2.Map2(convert)(maps.All(input)) {
		fmt.Printf("'%s' has length %d.\n", s, n)
	}

	// Unordered output:
	// 'abab' has length 4.
	// 'xyz' has length 3.
	// 'foofoofoo' has length 9.
}
