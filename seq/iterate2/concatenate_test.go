package iterate2_test

import (
	"slices"
	"testing"

	"github.com/GodsBoss/g/seq/iterate2"
)

func TestConcatenatingInfiniteSequence(t *testing.T) {
	numbers := func(yield func(index int, n int) bool) {
		n := 1
		for yield(n-1, n) {
			n++
		}
	}

	inputs := []int{-5, -10, -15}

	outputs := make([]int, 0)

	for _, n := range iterate2.Concatenate(numbers, slices.All(inputs)) {
		if n > 3 {
			break
		}

		outputs = append(outputs, n)
	}

	if len(outputs) != 3 {
		t.Fatalf("expected 3 outputs, got %+v", outputs)
	}

	if outputs[0] != 1 || outputs[1] != 2 || outputs[2] != 3 {
		t.Errorf("expected ouputs [1 2 3], got %+v", outputs)
	}
}
