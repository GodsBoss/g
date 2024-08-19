package bitseq_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/GodsBoss/g/seq/bitseq"
)

func TestAndInt8(t *testing.T) {
	testcases := map[string]testcase{
		"AND int8 default": bitTestcase[int8]{
			reduce:   bitseq.And[int8],
			inputs:   []int8{},
			expected: -1,
			bitSize:  8,
		},
		"AND uint8 with values": bitTestcase[uint8]{
			reduce:   bitseq.And[uint8],
			inputs:   []uint8{0b01011011, 0b11010001},
			expected: 0b01010001,
			bitSize:  8,
		},
		"AND uint8 reaching 0": bitTestcase[uint8]{
			reduce:   bitseq.And[uint8],
			inputs:   []uint8{0b10101100, 0b01010011},
			expected: 0b00000000,
			bitSize:  8,
		},
		"OR int8 default": bitTestcase[int8]{
			reduce:   bitseq.Or[int8],
			inputs:   []int8{},
			expected: 0,
			bitSize:  8,
		},
		"OR uint8 with values": bitTestcase[uint8]{
			reduce:   bitseq.Or[uint8],
			inputs:   []uint8{0b01101010, 0b00110011},
			expected: 0b01111011,
			bitSize:  8,
		},
		"OR uint8 reaching 0b11111111": bitTestcase[uint8]{
			reduce:   bitseq.Or[uint8],
			inputs:   []uint8{0b10101100, 0b01010011},
			expected: 0b11111111,
			bitSize:  8,
		},
	}

	for name := range testcases {
		t.Run(name, testcases[name].run)
	}
}

type testcase interface {
	run(t *testing.T)
}

type bitTestcase[T bitseq.Integer] struct {
	reduce   func(iter.Seq[T]) T
	inputs   []T
	expected T
	bitSize  int
}

func (tc bitTestcase[T]) run(t *testing.T) {
	actual := tc.reduce(slices.Values(tc.inputs))

	if actual != tc.expected {
		t.Errorf("\nExpected: %s\nActual  : %s", toBinaryString(tc.expected, tc.bitSize), toBinaryString(actual, tc.bitSize))
	}
}

func toBinaryString[T bitseq.Integer](value T, bitSize int) string {
	s := ""
	for i := 0; i < bitSize; i++ {
		if value&(1<<i) != 0 {
			s = "1" + s
		} else {
			s = "0" + s
		}
	}
	return s
}
