package iterate_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleToSequence() {
	ch := make(chan string, 3)
	ch <- "This"
	ch <- "is"
	ch <- "Sparta!"
	close(ch)

	for line := range iterate.FromChannel(ch) {
		fmt.Println(line)
	}

	// Output:
	// This
	// is
	// Sparta!
}

func ExampleFromSequence_exhaustion() {
	ch, cancel := iterate.IntoChannel(iterate.WithoutKeys(slices.All([]string{"This", "is", "Sparta!"})))
	defer cancel()

	for line := range ch {
		fmt.Println(line)
	}

	// Output:
	// This
	// is
	// Sparta!
}

func ExampleFromSequence_cancel() {
	numbers := func(yield func(int) bool) {
		n := 1
		for yield(n) {
			n++
		}
	}

	ch, cancel := iterate.IntoChannel(numbers)

	for n := range ch {
		if n < 5 {
			fmt.Printf("Current number is %d.\n", n)
		} else {
			cancel()
		}
	}

	// Output:
	// Current number is 1.
	// Current number is 2.
	// Current number is 3.
	// Current number is 4.
}
