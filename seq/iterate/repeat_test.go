package iterate_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/iterate"
)

func TestRepeatDoesNotCallWrappedIteratorOnZeroCount(t *testing.T) {
	called := false

	iterator := func(yield func(value struct{}) bool) {
		called = true
	}

	for _ = range iterate.Repeat[struct{}](iterate.Never, iterate.IsReusable)(iterator) {
	}

	if called {
		t.Errorf("expected wrapped iterator not to be called")
	}
}

func TestRepeatDoesNotRepeatEmptyIterators(t *testing.T) {
	iterator := func(yield func(value struct{}) bool) {}

	for _ = range iterate.Repeat[struct{}](iterate.InfiniteTimes, iterate.IsReusable)(iterator) {
	}
}
