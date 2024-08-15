package seq_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq"
)

func ExampleZip() {
	givenNames := []string{"Peter", "Paul", "Mary"}
	familyNames := []string{"Smith", "Miller", "Doe"}

	zipped := seq.Zip(
		seq.WithoutKeys(slices.All(givenNames)),
		seq.WithoutKeys(slices.All(familyNames)),
	)

	for givenName, familyName := range zipped {
		fmt.Printf("%s %s\n", givenName, familyName)
	}

	// Output:
	// Peter Smith
	// Paul Miller
	// Mary Doe
}
