package problem_test

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/problem"
)

func ExampleSetExtension() {
	details := problem.Details{
		Type:     "https://example.com/probs/out-of-credit",
		Title:    "You do not have enough credit.",
		Detail:   "Your current balance is 30, but that costs 50.",
		Instance: "/account/12345/msgs/abc",
	}

	type Extra struct {
		Balance  int      `json:"balance"`
		Accounts []string `json:"accounts"`
	}

	extra := Extra{
		Balance: 30,
		Accounts: []string{
			"/account/12345",
			"/account/67890",
		},
	}

	if err := problem.SetExtension(&details, extra); err != nil {
		panic(err)
	}

	data, err := json.Marshal(details)
	if err != nil {
		panic(err)
	}

	fields := make(map[string]any)

	if err := json.Unmarshal(data, &fields); err != nil {
		panic(err)
	}

	fmt.Printf("Balance: %v\n", fields["balance"])
	fmt.Printf("Accounts: %v\n", fields["accounts"])

	// Output:
	// Balance: 30
	// Accounts: [/account/12345 /account/67890]
}
