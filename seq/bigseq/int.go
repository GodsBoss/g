package bigseq

import (
	"iter"
	"math/big"
)

// IntSum produces the sum of the sequence's integers.
func IntSum(sequence iter.Seq[*big.Int]) *big.Int {
	result := new(big.Int)

	for n := range sequence {
		result = result.Add(result, n)
	}

	return result
}
