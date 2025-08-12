package csvseq_test

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/GodsBoss/g/seq/csvseq"
)

func TestEmpty(t *testing.T) {
	iterator := csvseq.Iterate(csv.NewReader(strings.NewReader("")))

	for record, err := range iterator {
		t.Errorf("expected no iteration, got %v; %v", record, err)
	}
}
