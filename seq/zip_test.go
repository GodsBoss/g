package seq_test

import (
	"testing"

	"github.com/GodsBoss/g/seq"
)

func TestZipCanBeStopped(t *testing.T) {
	numbers := func(yield func(int) bool) {
		n := 1
		for yield(n) {
			n++
		}
	}

	for i, j := range seq.Zip(numbers, numbers) {
		if i+j > 10 {
			break
		}
	}
}
