package iterate_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/iterate"
)

func TestToSequenceLeavesItemsInChannel(t *testing.T) {
	t.Parallel()

	ch := make(chan string, 5)

	ch <- "first"
	ch <- "second"
	ch <- "third"
	ch <- "fourth"
	ch <- "fifth"

	for value := range iterate.FromChannel(ch) {
		if value == "third" {
			break
		}
	}

	if len(ch) != 2 {
		t.Errorf("expected 2 items in channel, got %d", len(ch))
	}
}
