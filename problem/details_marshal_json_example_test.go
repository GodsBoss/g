package problem_test

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/problem"
)

func ExampleDetails_MarshalJSON() {
	details := problem.Details{
		Title:  "invalid fields",
		Status: 400,
		ExtensionMembers: map[string]any{
			"detail": map[string]any{
				"text": "There were invalid fields",
			},
			"fields": []string{"money", "age"},
		},
	}

	data, _ := json.MarshalIndent(details, "", " ")
	fmt.Printf("%s\n", data)

	// Output:
	// {
	//  "fields": [
	//   "money",
	//   "age"
	//  ],
	//  "status": 400,
	//  "title": "invalid fields",
	//  "type": "about:blank"
	// }
}
