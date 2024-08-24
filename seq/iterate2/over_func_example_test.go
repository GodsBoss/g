package iterate2_test

import (
	"fmt"
	"strconv"

	"github.com/GodsBoss/g/seq/iterate2"
)

func ExampleOverFunc() {
	var current int = 0
	counter := func() (int, string) {
		current++
		return current, strconv.Itoa(current)
	}

	for intValue, stringValue := range iterate2.OverFunc(counter) {
		if intValue > 3 {
			break
		}
		fmt.Println(intValue, stringValue)
	}

	// Output:
	// 1 1
	// 2 2
	// 3 3
}
