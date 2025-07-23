package bigseq_test

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/GodsBoss/g/seq/bigseq"
	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleIntSum() {
	fmt.Println(
		bigseq.IntSum(
			slices.Values(
				[]*big.Int{
					big.NewInt(7),
					big.NewInt(-3),
					big.NewInt(22),
				},
			),
		),
	)

	// Output:
	// 26
}

func ExampleIntSum_empty() {
	fmt.Println(bigseq.IntSum(iterate.Empty[*big.Int]))

	// Output:
	// 0
}

func ExampleIntProduct() {
	fmt.Println(
		bigseq.IntProduct(
			slices.Values(
				[]*big.Int{
					big.NewInt(8),
					big.NewInt(-7),
					big.NewInt(-2),
				},
			),
		),
	)

	// Output:
	// 112
}
