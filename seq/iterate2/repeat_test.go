package iterate2_test

import (
	"testing"

	"github.com/GodsBoss/g/seq/iterate2"
)

func TestRepeatDoesNotCallWrappedIteratorOnZeroCount(t *testing.T) {
	called := false

	iterator := func(yield func(first struct{}, second struct{}) bool) {
		called = true
	}

	for _ = range iterate2.Repeat[struct{}, struct{}](iterate2.Never, iterate2.IsReusable)(iterator) {
	}

	if called {
		t.Errorf("expected wrapped iterator not to be called")
	}
}

func TestRepeatDoesNotRepeatEmptyIterators(t *testing.T) {
	iterator := func(yield func(first struct{}, second struct{}) bool) {}

	for _ = range iterate2.Repeat[struct{}, struct{}](iterate2.InfiniteTimes, iterate2.IsReusable)(iterator) {
	}
}
