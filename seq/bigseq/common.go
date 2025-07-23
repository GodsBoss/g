package bigseq

import (
	"iter"

	"github.com/GodsBoss/g/seq/iterate"
)

func sum[Ptr interface {
	*T
	addable[Ptr]
}, T any](sequence iter.Seq[Ptr]) Ptr {
	var result Ptr = new(T)
	result = result.SetInt64(0)

	result = iterate.Reduce(result, result.Add)(sequence)

	return result
}

type addable[T any] interface {
	Add(T, T) T
	SetInt64(int64) T
}

func product[Ptr interface {
	*T
	multiplyable[Ptr]
}, T any](sequence iter.Seq[Ptr]) Ptr {
	var result Ptr = new(T)
	result = result.SetInt64(1)

	result = iterate.Reduce(result, result.Mul)(sequence)

	return result
}

type multiplyable[T any] interface {
	Mul(T, T) T
	SetInt64(int64) T
}
