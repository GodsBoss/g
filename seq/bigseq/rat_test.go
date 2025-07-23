package bigseq_test

import (
	"math/big"
	"testing"

	"github.com/GodsBoss/g/seq/bigseq"
	"github.com/GodsBoss/g/seq/iterate"
)

func TestRatProductOfEmptySequenceIsOne(t *testing.T) {
	one := big.NewRat(1, 1)
	product := bigseq.RatProduct(iterate.Empty[*big.Rat])

	if product.Cmp(one) != 0 {
		t.Errorf("expected %s to equal %s", product, one)
	}
}
