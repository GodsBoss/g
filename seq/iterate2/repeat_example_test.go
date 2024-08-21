package iterate2_test

import (
	"fmt"
	"slices"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleRepeat_reusable() {
	iterator := func(yield func(name string, age int) bool) {
		yield("Peter", 25)
		yield("Paul", 17)
		yield("Mary", 22)
	}

	for name, age := range iterate2.Repeat[string, int](2, iterate2.IsReusable)(iterator) {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Output:
	// Peter is 25 years old.
	// Paul is 17 years old.
	// Mary is 22 years old.
	// Peter is 25 years old.
	// Paul is 17 years old.
	// Mary is 22 years old.
}

func ExampleRepeat_notReusable() {
	var exhausted bool

	iterator := func(yield func(number int, name string) bool) {
		if exhausted {
			return
		}

		yield(1, "one")
		yield(2, "two")
		yield(3, "three")

		exhausted = true
	}

	for n, s := range iterate2.Repeat[int, string](2, iterate2.IsNotReusable)(iterator) {
		fmt.Printf("%d -> %s\n", n, s)
	}

	// Output:
	// 1 -> one
	// 2 -> two
	// 3 -> three
	// 1 -> one
	// 2 -> two
	// 3 -> three
}

func ExampleRepeat_infinite() {
	count := 0
	for _, s := range iterate2.Repeat[int, string](iterate2.InfiniteTimes, iterate2.IsReusable)(slices.All([]string{"xyz"})) {
		fmt.Println(s)
		count++
		if count >= 3 {
			break
		}
	}

	// Output:
	// xyz
	// xyz
	// xyz
}
