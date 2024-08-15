package channel_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq/channel"
)

func ExampleToSequence() {
	ch := make(chan string, 3)
	ch <- "This"
	ch <- "is"
	ch <- "Sparta!"
	close(ch)

	for line := range channel.ToSequence(ch) {
		fmt.Println(line)
	}

	// Output:
	// This
	// is
	// Sparta!
}
