package optional_test

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/optional"
)

func ExampleJSONValue_MarshalJSON_empty() {
	var value optional.JSONValue[int]

	data, err := json.Marshal(
		map[string]interface{}{
			"value": value,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)

	// Output:
	// {"value":null}
}

func ExampleJSONValue_MarshalJSON_nonempty() {
	value := optional.NewJSONValue(5000)

	data, err := json.Marshal(
		map[string]interface{}{
			"value": value,
		},
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)

	// Output:
	// {"value":5000}
}

func ExampleJSONValue_UnmarshalJSON_empty() {
	object := struct {
		Value optional.JSONValue[string]
	}{}

	err := json.Unmarshal([]byte(`{ "value": null }`), &object)

	if err != nil {
		panic(err)
	}

	optional.IfElse[string](
		object.Value,
		func(s string) {
			fmt.Println(s)
		},
		func() {
			fmt.Println("No value.")
		},
	)

	// Output:
	// No value.
}

func ExampleJSONValue_UnmarshalJSON_nonempty() {
	object := struct {
		Value optional.JSONValue[string]
	}{}

	err := json.Unmarshal(
		[]byte(`{ "value": "Hello, world!" }`),
		&object,
	)

	if err != nil {
		panic(err)
	}

	optional.IfElse[string](
		object.Value,
		func(s string) {
			fmt.Println(s)
		},
		func() {
			fmt.Println("No value.")
		},
	)

	// Output:
	// Hello, world!
}
