package iterate2_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/iterate2"
)

func TestEmpty(t *testing.T) {
	t.Parallel()

	called := false
	for range iterate2.Empty[struct{}, struct{}] {
		called = true
	}

	if called {
		t.Errorf("expected empty for loop")
	}
}
