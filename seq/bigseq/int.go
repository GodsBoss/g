package bigseq

import (
	"iter"
	"math/big"
)

// IntSum produces the sum of the sequence's integers. If the sequence is empty, zero is returned.
func IntSum(sequence iter.Seq[*big.Int]) *big.Int {
	return sum(sequence)
}
