package problem

import (
	"encoding/json"
	"math"
	"net/http"
)

type Details[Extension any] struct {
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

	// invalidMembers stores members that contained unexpected types on unmarshalling.
	invalidMembers map[string]any

	extension    *Extension
	extensionErr error
}

func (d Details[_]) StatusText() string {
	return http.StatusText(d.Status)
}

// InvalidMembers returns members that match the defined members of the problem details RFC, but contained values
// with different types than expected, e.g. a "status" with type boolean or an "instance" with type object.
// Clients must not change the return value of this method.
func (d Details[_]) InvalidMembers() map[string]any {
	return d.invalidMembers
}

func (d Details[Extension]) Extension() (*Extension, error) {
	return d.extension, d.extensionErr
}

func (d *Details[Extension]) UnmarshalJSON(data []byte) error {
	var tmp map[string]any

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	invalidMembers := make(map[string]any)

	d.Title = toType[string](tmp, "title", invalidMembers)
	d.Detail = toType[string](tmp, "detail", invalidMembers)
	d.Instance = toType[string](tmp, "instance", invalidMembers)

	statusAsFloat64 := toType[float64](tmp, "status", invalidMembers)
	d.Status = int(statusAsFloat64)
	if statusAsFloat64 < math.MinInt {
		d.Status = 0
	}
	if statusAsFloat64 > math.MaxInt {
		d.Status = 0
	}

	d.Type = toType[string](tmp, "type", invalidMembers)
	if d.Type == "" {
		d.Type = "about:blank"
	}

	d.invalidMembers = invalidMembers

	for _, field := range memberFields {
		delete(tmp, field)
	}

	d.extension, d.extensionErr = unmarshalExtension[Extension](tmp)

	return nil
}

func unmarshalExtension[Extension any](remainingFields map[string]any) (*Extension, error) {
	jsonStringForExtension, _ := json.Marshal(remainingFields)
	var extension Extension
	if err := json.Unmarshal([]byte(jsonStringForExtension), &extension); err != nil {
		return nil, err
	}
	return &extension, nil
}

var memberFields = []string{
	"detail",
	"instance",
	"status",
	"title",
	"type",
}

func toType[T any](generic map[string]any, name string, otherwiseTarget map[string]any) (typed T) {
	value, ok := generic[name]
	if !ok {
		return
	}

	typed, ok = value.(T)
	if ok {
		return
	}

	otherwiseTarget[name] = value

	return
}
