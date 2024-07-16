package problem_test

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/problem"
)

func ExampleGetExtension() {
	type Extra struct {
		Balance  int
		Accounts []string
	}

	rawProblemDetails := `
		{
			"type": "https://example.com/probs/out-of-credit",
			"title": "You do not have enough credit.",
			"detail": "Your current balance is 30, but that costs 50.",
			"instance": "/account/12345/msgs/abc",
			"balance": 30,
			"accounts": ["/account/12345", "/account/67890"]
		}`

	var problemDetails problem.Details

	if err := json.Unmarshal([]byte(rawProblemDetails), &problemDetails); err != nil {
		panic(err)
	}

	extra, err := problem.GetExtension[Extra](problemDetails)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Balance: %d\n", extra.Balance)
	for i, account := range extra.Accounts {
		fmt.Printf("Account %d: %s\n", i+1, account)
	}

	// Output:
	// Balance: 30
	// Account 1: /account/12345
	// Account 2: /account/67890
}
