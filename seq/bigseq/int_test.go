package bigseq_test

import (
	"math/big"
	"testing"

	"github.com/GodsBoss/g/seq/bigseq"
	"github.com/GodsBoss/g/seq/iterate"
)

func TestIntProductOfEmptySequenceIsOne(t *testing.T) {
	one := big.NewInt(1)
	product := bigseq.IntProduct(iterate.Empty[*big.Int])

	if product.Cmp(one) != 0 {
		t.Errorf("expected %s to equal %s", product, one)
	}
}
