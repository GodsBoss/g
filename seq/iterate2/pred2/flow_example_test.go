package pred2_test

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/GodsBoss/g/seq/iterate2"
	"github.com/GodsBoss/g/seq/iterate2/pred2"
)

func ExampleContextIsValid() {
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel the context after 50 milliseconds. In practice a context would more likely be canceled
	// when a request times out or a server was signalled to be shutdown.
	go func() {
		time.Sleep(time.Millisecond * 50)
		cancel()
	}()

	items := []int{2, 3, 5, 7, 11, 13, 17, 19}

	for _, n := range iterate2.While(pred2.ContextIsValid[int, int](ctx))(slices.All(items)) {
		fmt.Println(n)
		time.Sleep(time.Millisecond * 15)
	}

	// Output:
	// 2
	// 3
	// 5
	// 7
}

func ExampleUntilCanceled() {
	items := []int{2, 3, 5, 7, 11, 13, 17, 19}

	untilCanceled, cancel := pred2.UntilCanceled[int, int]()

	// Cancel after 50 milliseconds.
	go func() {
		time.Sleep(time.Millisecond * 50)
		cancel()
	}()

	for _, n := range iterate2.While(untilCanceled)(slices.All(items)) {
		fmt.Println(n)
		time.Sleep(time.Millisecond * 15)
	}

	// Output:
	// 2
	// 3
	// 5
	// 7
}
