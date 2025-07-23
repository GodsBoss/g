// Package bigseq handles types of the math/big package, namely big.Int, big.Rat and big.Float,
// in an iterator way.
//
// Currently, only big.Int and big.Rat are supported as big.Float handling is a bit more complicated,
// due to its additional precision and rounding mode.
package bigseq
