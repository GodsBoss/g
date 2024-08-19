package boolseq_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/GodsBoss/g/seq/boolseq"
)

func TestBool(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		values []bool

		f func(iter.Seq[bool]) bool

		expected bool
	}{
		"AND []": {
			f:        boolseq.And,
			expected: true,
		},
		"AND [true, false]": {
			values:   []bool{true, false},
			f:        boolseq.And,
			expected: false,
		},
		"OR []": {
			f:        boolseq.Or,
			expected: false,
		},
		"OR [false, true]": {
			values:   []bool{false, true},
			f:        boolseq.Or,
			expected: true,
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				t.Parallel()

				actual := testcase.f(slices.Values(testcase.values))

				if actual != testcase.expected {
					t.Errorf("expected %t, got %t", testcase.expected, actual)
				}
			},
		)
	}
}
