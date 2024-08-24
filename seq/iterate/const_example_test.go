package iterate_test

import (
	"fmt"

	"github.com/GodsBoss/g/seq/iterate"
)

func ExampleConst() {
	count := 0
	for name := range iterate.Const("世界") {
		count++
		if count > 3 {
			break
		}
		fmt.Printf("Hello, %s!\n", name)
	}

	// Output:
	// Hello, 世界!
	// Hello, 世界!
	// Hello, 世界!
}
