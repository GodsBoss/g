// Package bigseq handles types of the math/big package in a Go 1.23 iterator way.
//
// Currently, only big.Int and big.Rat are supported as big.Float handling is a bit more complicated,
// due to its additional precision and rounding mode.
//
// Although Int.Or is supported, Int.And is not as there is no intuitive answer for an empty sequence.
// For finite integers, all bits are set to 1, but technically big.Int represents arbitrarely large
// integers, so there is no upper limit to the number of bits.
package bigseq
