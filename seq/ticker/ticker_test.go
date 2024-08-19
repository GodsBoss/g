package ticker_test

import (
	"testing"
	"time"

	"github.com/GodsBoss/g/seq/ticker"
)

func TestTimer(t *testing.T) {
	t.Parallel()

	end := time.Now().Add(time.Millisecond * 10)
	times := make([]time.Time, 0)

	for timestamp := range ticker.New(time.Millisecond * 3) {
		times = append(times, timestamp)
		if timestamp.After(end) {
			break
		}
	}

	if len(times) != 4 {
		t.Errorf("expected %d ticks, got %d", 4, len(times))
	}
}

func TestTimerSkipsSlowConsumers(t *testing.T) {
	t.Parallel()

	end := time.Now().Add(time.Millisecond * 10)
	times := make([]time.Time, 0)

	for timestamp := range ticker.New(time.Millisecond * 3) {
		time.Sleep(time.Millisecond * 10)
		times = append(times, timestamp)
		if timestamp.After(end) {
			break
		}
	}

	if len(times) != 3 {
		t.Errorf("expected %d ticks, got %d", 3, len(times))
	}
}
