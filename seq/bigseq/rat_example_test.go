package bigseq_test

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/GodsBoss/g/seq/bigseq"
	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleRatSum() {
	fmt.Println(
		bigseq.RatSum(
			slices.Values(
				[]*big.Rat{
					big.NewRat(2, 3),
					big.NewRat(-7, 5),
					big.NewRat(1, 8),
				},
			),
		),
	)

	// Output:
	// -73/120
}

func ExampleRatSum_empty() {
	fmt.Println(bigseq.RatSum(iterate.Empty[*big.Rat]))

	// Output:
	// 0/1
}

func ExampleRatProduct() {
	fmt.Println(
		bigseq.RatProduct(
			slices.Values(
				[]*big.Rat{
					big.NewRat(2, -3),
					big.NewRat(7, 13),
					big.NewRat(-39, 5),
				},
			),
		),
	)

	// Output:
	// 14/5
}

func ExampleRatProduct_empty() {
	fmt.Println(bigseq.RatProduct(iterate.Empty[*big.Rat]))

	// Output:
	// 1/1
}
