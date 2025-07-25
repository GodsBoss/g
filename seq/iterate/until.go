package iterate

import (
	"iter"

	"github.com/GodsBoss/g/seq/iterate/pred"
)

// Until yields values from a sequence until the given predicate becomes true.
//
// See the pred subpackage for pre-defined predicates.
//
// The resulting iterator is finite if the predicate ever becomes true.
//
// Reusable if underlying sequence is reusable.
func Until[Value any](predicate func(Value) bool) func(sequence iter.Seq[Value]) iter.Seq[Value] {
	return While(pred.Not(predicate))
}
