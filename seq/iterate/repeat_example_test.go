package iterate_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleRepeat_reusable() {
	for s := range iterate.Repeat[string](2, iterate.IsReusable)(slices.Values([]string{"one", "two", "three"})) {
		fmt.Println(s)
	}

	// Output:
	// one
	// two
	// three
	// one
	// two
	// three
}

func ExampleRepeat_notReusable() {
	ch := make(chan string, 3)
	ch <- "one"
	ch <- "two"
	ch <- "three"
	close(ch)

	for s := range iterate.Repeat[string](2, iterate.IsNotReusable)(iterate.FromChannel(ch)) {
		fmt.Println(s)
	}

	// Output:
	// one
	// two
	// three
	// one
	// two
	// three
}

func ExampleRepeat_infinite() {
	count := 0
	for s := range iterate.Repeat[string](iterate.InfiniteTimes, iterate.IsReusable)(slices.Values([]string{"xyz"})) {
		fmt.Println(s)
		count++
		if count >= 3 {
			break
		}
	}

	// Output:
	// xyz
	// xyz
	// xyz
}
