package bigseq

import (
	"iter"
	"math/big"
)

// IntSum produces the sum of the sequence's integers. If the sequence is empty, zero is returned.
func IntSum(sequence iter.Seq[*big.Int]) *big.Int {
	return sum(sequence)
}

// IntProduct produces the product of the sequence's integers. If the sequence is empty, one is returned.
func IntProduct(sequence iter.Seq[*big.Int]) *big.Int {
	return product(sequence)
}
