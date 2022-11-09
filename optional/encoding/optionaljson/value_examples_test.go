package optionaljson_test

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/optional"
	"github.com/GodsBoss/g/optional/encoding/optionaljson"
)

func ExampleValue_MarshalJSON_empty() {
	var value optionaljson.Value[int]

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

func ExampleValue_MarshalJSON_nonempty() {
	value := optionaljson.NewValue(5000)

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

func ExampleValue_UnmarshalJSON_empty() {
	object := struct {
		Value optionaljson.Value[string]
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

func ExampleValue_UnmarshalJSON_nonempty() {
	object := struct {
		Value optionaljson.Value[string]
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
