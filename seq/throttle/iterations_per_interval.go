package throttle

import (
	"errors"
	"fmt"
	"time"
)

// IterationsPerInterval is a blueprint for creating throttling strategies that pass at most a certain number of items within a given timeframe
// from a sequence.
type IterationsPerInterval struct {
	// Maximum is the maximum amount of iterations allowed within the timespan defined by Interval.
	// Must be > 0.
	Maximum int

	// Timeframe defines the timespan during which at most Maximum iterations may be invoked.
	// Must be > 0.
	Timeframe time.Duration
}

// Strategy creates a strategy with the configuration from the blueprint. The Strategy is independent from the blueprint afterwards.
func (strategyBlueprint IterationsPerInterval) Strategy() (Strategy, error) {
	var errs []error

	maximum, timeframe := strategyBlueprint.Maximum, strategyBlueprint.Timeframe

	if maximum <= 0 {
		errs = append(errs, fmt.Errorf("maximum iterations per interval must be > 0, not %d", maximum))
	}

	if timeframe <= 0 {
		errs = append(errs, fmt.Errorf("interval for maximum iterations per interval must be > 0, not %s", timeframe))
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	return iterationsPerInterval{
		maximum:   maximum,
		timeframe: timeframe,
	}, nil
}

type iterationsPerInterval struct {
	maximum   int
	timeframe time.Duration
}

func (strategy iterationsPerInterval) Waiter() Waiter {
	lastInvocations := make([]time.Time, 0)

	return WaiterFunc(
		func() {
			// Remove old invocations.
			for len(lastInvocations) > 0 && time.Since(lastInvocations[0]) >= strategy.timeframe {
				lastInvocations = lastInvocations[1:]
			}

			// If the invocation count is less than the allowed maximum we can immediately return.
			if len(lastInvocations) < strategy.maximum {
				lastInvocations = append(lastInvocations, time.Now())
				return
			}

			// Wait until the oldest invocation expires.
			time.Sleep(strategy.timeframe - time.Since(lastInvocations[0]))
			lastInvocations = append(lastInvocations[1:], time.Now())
		},
	)
}
