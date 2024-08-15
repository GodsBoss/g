package channel_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/channel"
)

func TestToSequenceLeavesItemsInChannel(t *testing.T) {
	ch := make(chan string, 5)

	ch <- "first"
	ch <- "second"
	ch <- "third"
	ch <- "fourth"
	ch <- "fifth"

	for value := range channel.ToSequence(ch) {
		if value == "third" {
			break
		}
	}

	if len(ch) != 2 {
		t.Errorf("expected 2 items in channel, got %d", len(ch))
	}
}
