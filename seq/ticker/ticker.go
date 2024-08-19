package ticker

import (
	"iter"
	"time"
)

// New creates an iterator backed by a ticker. Drops ticks for slow consumers, just as time.Ticker does.
func New(interval time.Duration) iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		t := time.Tick(interval)
		for yield(<-t) {
		}
	}
}
