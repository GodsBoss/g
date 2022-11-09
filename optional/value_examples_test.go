package optional_test

import (
	"fmt"

	"github.com/GodsBoss/g/optional"
)

func ExampleValue_empty() {
	value := optional.NewEmpty[int]()

	item, ok := optional.GetItem(value)
	fmt.Println(item)
	fmt.Println(ok)

	fmt.Println(optional.HasItem(value))
	fmt.Println(optional.IsEmpty(value))
	fmt.Println(optional.Len(value))
	fmt.Println(len(optional.ToSlice(value)))

	// Output:
	// 0
	// false
	// false
	// true
	// 0
	// 0
}

func ExampleValue_withItem() {
	value := optional.NewWithItem("Hello, world!")

	item, ok := optional.GetItem(value)
	fmt.Println(item)
	fmt.Println(ok)

	fmt.Println(optional.HasItem(value))
	fmt.Println(optional.IsEmpty(value))
	fmt.Println(optional.Len(value))
	fmt.Println(len(optional.ToSlice(value)))

	// Output:
	// Hello, world!
	// true
	// true
	// false
	// 1
	// 1
}

func ExampleIfElse() {
	valueWithItem := optional.NewWithItem("I am an item.")
	emptyValue := optional.NewEmpty[string]()

	printItem := func(s string) {
		fmt.Println(s)
	}
	printDefault := func() {
		fmt.Println("I am empty.")
	}

	optional.IfElse(
		valueWithItem,
		printItem,
		printDefault,
	)

	optional.IfElse(
		emptyValue,
		printItem,
		printDefault,
	)

	// Output:
	// I am an item.
	// I am empty.
}

func ExampleFromPointer_nil() {
	var ptr *int

	value := optional.FromPointer(ptr)

	fmt.Println(optional.HasItem(value))

	// Output:
	// false
}

func ExampleFromPointer_nonnil() {
	s := "Hello!"

	value := optional.FromPointer(&s)

	if item, ok := optional.GetItem(value); ok {
		fmt.Println(*item)
	}

	// Output:
	// Hello!
}

func ExampleMapItem() {
	messages := map[string]string{
		"departure": "Goodbye!",
	}

	optional.
		MapItem(messages, "departure").
		Invoke(
			func(s string) {
				fmt.Println(s)
			},
		)

	fmt.Println(
		optional.IsEmpty(
			optional.MapItem(messages, "arrival"),
		),
	)

	// Output:
	// Goodbye!
	// true
}
