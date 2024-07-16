package problem_test

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/problem"
)

func ExampleDetails_UnmarshalJSON() {
	input := `{
		"type": "https://example.com/problems/666",
		"detail": {
			"lang": "de-de",
			"text": "Feldwertvalidierungsfehler"
		},
		"extra": [
			"foo",
			"bar"
		]
	}`

	var details problem.Details

	if err := json.Unmarshal([]byte(input), &details); err != nil {
		panic(err)
	}

	fmt.Printf("Type   : %s\n", details.Type)
	fmt.Printf("Extra  : %v\n", details.ExtensionMembers["extra"])
	fmt.Printf("Invalid: %v\n", details.InvalidMembers["detail"])

	// Output:
	// Type   : https://example.com/problems/666
	// Extra  : [foo bar]
	// Invalid: map[lang:de-de text:Feldwertvalidierungsfehler]
}
