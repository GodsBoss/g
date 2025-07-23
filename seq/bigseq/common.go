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
