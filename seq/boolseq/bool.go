// Package boolseq contains helpers for sequences of booleans.
package boolseq

import "iter"

// And applies a logical AND to all items of a sequence. Returns false if it encounters any false, else true.
// Stops pulling from sequence as soon as the first false is found.
func And(sequence iter.Seq[bool]) bool {
	for b := range sequence {
		if !b {
			return false
		}
	}

	return true
}

// Or applies a logical OR to all items of a sequence. Returns true if it encounters any true, else false.
// Stops pulling from sequence as soon as the first true is found.
func Or(sequence iter.Seq[bool]) bool {
	for b := range sequence {
		if b {
			return true
		}
	}

	return false
}
