package seq_test

import (
	"fmt"
	"maps"

	"github.com/GodsBoss/g/seq"
)

func ExampleWithoutValues() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range seq.WithoutValues(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// Peter
	// Paul
	// Mary
}

func ExampleWithoutKeys() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for name := range seq.WithoutKeys(maps.All(ages)) {
		fmt.Println(name)
	}

	// Unordered output:
	// 33
	// 22
	// 27
}

func ExampleSwapKeyValues() {
	ages := map[string]int{
		"Peter": 33,
		"Paul":  22,
		"Mary":  27,
	}

	for age, name := range seq.SwapKeyValues(maps.All(ages)) {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Unordered output:
	// Peter is 33 years old.
	// Paul is 22 years old.
	// Mary is 27 years old.
}
