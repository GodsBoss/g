package bigseq

import (
	"iter"
	"math/big"
)

// RatSum produces the sum of the sequence's rational numbers. If the sequence is empty, zero is returned.
func RatSum(sequence iter.Seq[*big.Rat]) *big.Rat {
	return sum(sequence)
}
