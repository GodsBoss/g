package bigseq

import "iter"

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

func product[Ptr interface {
	*T
	multiplyable[Ptr]
}, T any](sequence iter.Seq[Ptr]) Ptr {
	var result Ptr = new(T)
	result = result.SetInt64(1)

	for n := range sequence {
		result = result.Mul(result, n)
	}

	return result
}

type multiplyable[T any] interface {
	Mul(T, T) T
	SetInt64(int64) T
}
