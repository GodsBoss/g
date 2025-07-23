package bigseq

import (
	"iter"
	"math/big"
)

// RatSum produces the sum of the sequence's rational numbers.
func RatSum(sequence iter.Seq[*big.Rat]) *big.Rat {
	return sum(sequence)
}
