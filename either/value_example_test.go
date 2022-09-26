package either_test

import (
	"fmt"

	"github.com/GodsBoss/g/either"
)

func ExampleInvoke() {
	value := either.NewFirst[int, int](0)

	either.Invoke(value, nil, nil)
}

func ExampleSwap() {
	value := either.NewFirst[int, string](12345)

	swapped := either.Swap(value)

	if item, ok := either.GetSecond(swapped); ok {
		fmt.Printf("Second: %d\n", item)
	}

	swappedAgain := either.Swap(swapped)

	if item, ok := either.GetFirst(swappedAgain); ok {
		fmt.Printf("First: %d\n", item)
	}

	// Output:
	// Second: 12345
	// First: 12345
}

func ExampleToSlices_first() {
	value := either.NewFirst[int, string](1981)

	numbers, messages := either.ToSlices(value)

	fmt.Printf("number count  : %d\n", len(numbers))
	fmt.Printf("messages count: %d\n", len(messages))

	fmt.Println(numbers[0])

	// Output:
	// number count  : 1
	// messages count: 0
	// 1981
}

func ExampleToSlices_second() {
	value := either.NewSecond[int, string]("1981")

	numbers, messages := either.ToSlices(value)

	fmt.Printf("number count  : %d\n", len(numbers))
	fmt.Printf("messages count: %d\n", len(messages))

	fmt.Println(messages[0])

	// Output:
	// number count  : 0
	// messages count: 1
	// 1981
}
