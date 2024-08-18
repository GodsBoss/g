package mathseq_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/mathseq"
)

func ExampleSum() {
	numbers := []int{2, 3, 5, 7}
	sum := mathseq.Sum(slices.Values(numbers))
	fmt.Println(sum)

	// Output:
	// 17
}

func ExampleProduct() {
	numbers := []int{2, 3, 5, 7}
	product := mathseq.Product(slices.Values(numbers))
	fmt.Println(product)

	// Output:
	// 210
}
