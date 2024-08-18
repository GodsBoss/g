package mathseq

import (
	"cmp"
	"iter"
)

// Max returns the maximum of items in the sequence, if any. When the sequence is empty, ok is false, it's true otherwise.
// For floating-point Item, Max propagates NaNs (any NaN value i x forces the output to be NaN). This is the same behaviour as slices.Max.
func Max[Item cmp.Ordered](items iter.Seq[Item]) (result Item, ok bool) {
	return MaxFunc[Item](cmp.Compare)(items)

}

// MaxFunc returns the maximal value from a sequence, using cmp to compare elements. When the sequence is empty, ok is false, it's true otherwise.
// If there is more than one maximal element according to the cmp function, MaxFunc returns the first one. This is the same behaviour as slices.MaxFunc.
func MaxFunc[Item cmp.Ordered](cmp func(current, next Item) int) func(iter.Seq[Item]) (result Item, ok bool) {
	return func(sequence iter.Seq[Item]) (result Item, ok bool) {
		for current := range sequence {
			if !ok {
				result = current
				ok = true
				continue
			}

			if cmp(result, current) < 0 {
				result = current
			}

		}

		return
	}
}

// Min returns the minimum of items in the sequence, if any. When the sequence is empty, ok is false, it's true otherwise.
// For floating-point Item, Min propagates NaNs (any NaN value i x forces the output to be NaN). This is the same behaviour as slices.Min.
func Min[Item cmp.Ordered](items iter.Seq[Item]) (result Item, ok bool) {
	return MinFunc[Item](cmp.Compare)(items)

}

// MinFunc returns the minimal value from a sequence, using cmp to compare elements. When the sequence is empty, ok is false, it's true otherwise.
// If there is more than one minimal element according to the cmp function, MinFunc returns the first one. This is the same behaviour as slices.MinFunc.
func MinFunc[Item cmp.Ordered](cmp func(current, next Item) int) func(iter.Seq[Item]) (result Item, ok bool) {
	return func(sequence iter.Seq[Item]) (result Item, ok bool) {
		for current := range sequence {
			if !ok {
				result = current
				ok = true
				continue
			}

			if cmp(result, current) > 0 {
				result = current
			}
		}

		return
	}
}
