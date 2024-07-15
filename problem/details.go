package problem

import (
	"encoding/json"
	"math"
	"net/http"
)

type Details struct {
	// Type is a URI reference that identifies the problem type. Defaults to "about:blank".
	Type string

	// Status is the HTTP status code. May be 0.
	Status int

	// Title is a human-readable summary of the problem type.
	Title string

	// Detail is a human-readable explanation specific to this occurence of the problem.
	Detail string

	// Instance is a URI reference that identifies the specific occurence of the problem. Optional.
	Instance string
}

func (d Details) StatusText() string {
	return http.StatusText(d.Status)
}

func (d *Details) UnmarshalJSON(data []byte) error {
	var tmp map[string]any

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	d.Type = toType[string](tmp["type"])
	d.Title = toType[string](tmp["title"])
	d.Detail = toType[string](tmp["detail"])
	d.Instance = toType[string](tmp["instance"])

	statusAsFloat64 := toType[float64](tmp["status"])
	d.Status = int(statusAsFloat64)
	if statusAsFloat64 < math.MinInt || statusAsFloat64 > math.MaxInt {
		d.Status = 0
	}

	if d.Type == "" {
		d.Type = "about:blank"
	}

	return nil
}

func toType[T any](value any) T {
	typed, ok := value.(T)
	if ok {
		return typed
	}

	var zero T

	return zero
}
