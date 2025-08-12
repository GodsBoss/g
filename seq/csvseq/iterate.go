package csvseq

import (
	"encoding/csv"
	"io"
	"iter"
)

// Iterate creates an iterator over the records of the CSV reader. Errors occuring might be csv.ErrFieldCount or
// parse errors. As Iterate is supposed to consume all the records, the error is never io.EOF.
func Iterate(reader *csv.Reader) iter.Seq2[[]string, error] {
	return func(yield func([]string, error) bool) {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				return
			}

			if !yield(record, err) {
				return
			}
		}
	}
}
