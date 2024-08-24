package iterate_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/iterate"
)

func TestEmpty(t *testing.T) {
	t.Parallel()

	called := false
	for range iterate.Empty[struct{}] {
		called = true
	}

	if called {
		t.Errorf("expected empty for loop")
	}
}
