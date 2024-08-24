package iterate2_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleConst() {
	count := 0
	for name, age := range iterate2.Const("Someone", 69) {
		count++
		if count > 3 {
			break
		}
		fmt.Printf("%s is %d years old. Nice!\n", name, age)
	}

	// Output:
	// Someone is 69 years old. Nice!
	// Someone is 69 years old. Nice!
	// Someone is 69 years old. Nice!
}
