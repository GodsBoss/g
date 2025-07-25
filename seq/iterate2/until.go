package iterate2

import (
	"iter"

	"github.com/GodsBoss/g/seq/iterate2/pred2"
)

// Until yields values from a sequence until the given predicate becomes true.
//
// See the pred2 subpackage for pre-defined predicates.
//
// The resulting iterator is finite if the predicate ever becomes true.
//
// Reusable if underlying sequence is reusable.
func Until[First any, Second any](predicate func(First, Second) bool) func(iter.Seq2[First, Second]) iter.Seq2[First, Second] {
	return While(pred2.Not(predicate))
}
