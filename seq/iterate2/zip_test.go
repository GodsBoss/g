package iterate2_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/iterate2"
)

func TestZipCanBeStopped(t *testing.T) {
	numbers := func(yield func(int) bool) {
		n := 1
		for yield(n) {
			n++
		}
	}

	for i, j := range iterate2.Zip(numbers, numbers) {
		if i+j > 10 {
			break
		}
	}
}
