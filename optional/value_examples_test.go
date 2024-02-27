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

func ExampleIf() {
	valueWithItem := optional.NewWithItem("I am an item.")
	emptyValue := optional.NewEmpty[string]()

	optional.If(
		valueWithItem,
		func(s string) {
			fmt.Println("First: " + s)
		},
	)

	optional.If(
		emptyValue,
		func(s string) {
			fmt.Println("Second: " + s)
		},
	)

	// Output:
	// First: I am an item.
}

func ExampleElse() {
	valueWithItem := optional.NewWithItem("I am an item.")
	emptyValue := optional.NewEmpty[string]()

	optional.Else(
		valueWithItem,
		func() {
			fmt.Println("First.")
		},
	)

	optional.Else(
		emptyValue,
		func() {
			fmt.Println("Second.")
		},
	)

	// Output:
	// Second.
}

func ExamplePointer_nil() {
	var ptr *int

	value := optional.Pointer(ptr)

	fmt.Println(optional.HasItem(value))

	// Output:
	// false
}

func ExamplePointer_nonnil() {
	s := "Hello!"

	value := optional.Pointer(&s)

	if item, ok := optional.GetItem(value); ok {
		fmt.Println(*item)
	}

	// Output:
	// Hello!
}

func ExampleDereference_nil() {
	var ptr *int

	value := optional.Dereference(ptr)

	fmt.Println(optional.HasItem(value))

	// Output:
	// false
}

func ExampleDereference_nonnil() {
	s := "Hello!"

	value := optional.Dereference(&s)

	if item, ok := optional.GetItem(value); ok {
		fmt.Println(item)
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

func ExampleSliceItem() {
	items := []string{
		"foo",
		"bar",
		"baz",
	}

	optional.
		SliceItem(items, 1).
		Invoke(
			func(s string) {
				fmt.Println(s)
			},
		)

	fmt.Println(
		optional.IsEmpty(
			optional.SliceItem(items, -1),
		),
	)

	fmt.Println(
		optional.IsEmpty(
			optional.SliceItem(items, 8),
		),
	)

	// Output:
	// bar
	// true
	// true
}

func ExampleFromOKResult() {
	// f stands in for functions that return a value and a boolean to indicate whether the
	// return value is meaningful, e.g. os.LookupEnv or strings.CutPrefix.
	f := func(s string, ok bool) (string, bool) {
		return s, ok
	}

	optional.
		FromOKResult(f("Hello, world!", true)).
		Invoke(
			func(s string) {
				fmt.Println(s)
			},
		)

	optional.
		FromOKResult(f("none", false)).
		Invoke(
			func(s string) {
				fmt.Println(s)
			},
		)

	// Output:
	// Hello, world!
}
