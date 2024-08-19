package mathseq_test

import (
	"slices"
	"testing"

	"github.com/GodsBoss/g/seq/mathseq"
)

func TestMinWithoutValues(t *testing.T) {
	t.Parallel()

	_, ok := mathseq.Min(slices.Values(make([]float64, 0)))

	if ok {
		t.Errorf("expected ok to be false.")
	}
}

func TestMaxWithoutValues(t *testing.T) {
	t.Parallel()

	_, ok := mathseq.Max(slices.Values(make([]float64, 0)))

	if ok {
		t.Errorf("expected ok to be false.")
	}
}
