package throttle_test

import (
	"testing"
	"time"

	"github.com/GodsBoss/g/seq/throttle"
)

func TestInvalidIterationsPerIntervalBlueprint(t *testing.T) {
	t.Parallel()

	testcases := map[string]throttle.IterationsPerInterval{
		"maximum too low": throttle.IterationsPerInterval{
			Maximum:   -1,
			Timeframe: time.Second,
		},
		"timeframe too low": throttle.IterationsPerInterval{
			Maximum:   5,
			Timeframe: -time.Second,
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				t.Parallel()

				strategy, err := testcase.Strategy()
				if err == nil {
					t.Error("expected an error")
				}
				if strategy != nil {
					t.Error("expected no strategy")
				}
			},
		)
	}
}

func TestIterationsPerInterval(t *testing.T) {
	t.Parallel()

	iterations := 0
	stop := make(chan struct{})
	strategyErr := make(chan error)

	go func() {
		strategy, err := throttle.IterationsPerInterval{
			Maximum:   5,
			Timeframe: time.Millisecond * 3,
		}.Strategy()

		if err != nil {
			strategyErr <- err
			return
		}

		iterator := func(yield func(struct{}) bool) {
			for {
				select {
				case <-stop:
					return
				default:
					if !yield(struct{}{}) {
						return
					}
				}
			}
		}

		iterator = throttle.Iteration[struct{}](strategy)(iterator)

		for _ = range iterator {
			iterations++
		}
	}()

	waitTime := time.Tick(time.Millisecond * 10)

	select {
	case err := <-strategyErr:
		t.Fatalf("could not create iterations per interval strategy: %+v", err)
	case <-waitTime:
	}

	close(stop)

	if iterations != 20 {
		t.Errorf("expected %d iterations, got %d", 20, iterations)
	}
}
