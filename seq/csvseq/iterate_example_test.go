package csvseq_test

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strings"

	"github.com/GodsBoss/g/seq/csvseq"
)

func ExampleIterate_tooFewFields() {
	raw := strings.Join(
		[]string{
			"Foo,Bar,Baz",
			"a,b",
			"1,2,3",
		},
		"\n",
	)

	reader := csv.NewReader(strings.NewReader(raw))
	reader.FieldsPerRecord = 3

	for record, err := range csvseq.Iterate(reader) {
		if err != nil {
			fmt.Println("Invalid record " + strings.Join(record, ",") + " - " + err.Error())
			break
		}
	}

	// Output:
	// Invalid record a,b - record on line 2: wrong number of fields
}

func ExampleIterate_exhaustion() {
	raw := strings.Join(
		[]string{
			"1,2,3",
			"a,b,c",
			"x,y,z",
		},
		"\n",
	)

	reader := csv.NewReader(strings.NewReader(raw))

	var errs error

	for record, err := range csvseq.Iterate(reader) {
		fmt.Println(record[0])
		errs = errors.Join(errs, err)
	}

	if errs != nil {
		fmt.Println(errs)
	}

	// Output:
	// 1
	// a
	// x
}
