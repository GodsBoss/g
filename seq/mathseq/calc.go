package mathseq

import (
	"iter"
)

// Numeric is a constraint that contains all numeric types.
type Numeric interface {
	float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Sum returns the sum of all the numbers from the sequence. If the sequence is empty, 0 is returned.
func Sum[Number Numeric](sequence iter.Seq[Number]) Number {
	var sum Number = 0

	for summand := range sequence {
		sum += summand
	}

	return sum
}

// Product returns the product of all the numbers from the sequence. If the sequence is empty, 1 is returned.
func Product[Number Numeric](sequence iter.Seq[Number]) Number {
	var product Number = 1

	for factor := range sequence {
		product *= factor
	}

	return product
}
