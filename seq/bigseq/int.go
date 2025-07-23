package bigseq

import (
	"iter"
	"math/big"
)

// IntSum produces the sum of the sequence's integers.
func IntSum(sequence iter.Seq[*big.Int]) *big.Int {
	return sum[*big.Int](sequence)
}

func sum[Ptr interface {
	*T
	addable[Ptr]
}, T any](sequence iter.Seq[Ptr]) Ptr {
	var result Ptr = new(T)

	for n := range sequence {
		result = result.Add(result, n)
	}

	return result
}

type addable[T any] interface {
	Add(T, T) T
}
