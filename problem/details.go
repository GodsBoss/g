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

	// InvalidMembers contains fields that matched the problem details member names, but not their types.
	// Populated when unmarshalling, ignored on marshalling.
	InvalidMembers map[string]any

	// ExtensionMembers contains the extension members of this problem details instance.
	//
	// Populated with fields that are not defined by the problem detail RfC when unmarshalling.
	//
	// When marshalling, these fields and the members of the problem details instance are merged.
	// Fields that would result in the same field names as those of the problem details are ignored,
	// even if the corresponding fields contain zero values, e.g. even when Title is not set,
	// an extension member "title" would not be marshalled.
	ExtensionMembers map[string]any
}

func (d Details) StatusText() string {
	return http.StatusText(d.Status)
}

func (d Details) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	for k, v := range d.ExtensionMembers {
		m[k] = v
	}

	for _, k := range memberFields {
		delete(m, k)
	}

	if d.Type != "" {
		m["type"] = d.Type
	}

	if d.Status != 0 {
		m["status"] = d.Status
	}

	if d.Title != "" {
		m["title"] = d.Title
	}

	if d.Detail != "" {
		m["detail"] = d.Detail
	}

	if d.Instance != "" {
		m["instance"] = d.Instance
	}
	if d.Type == "" {
		m["type"] = "about:blank"
	}

	return json.Marshal(m)
}

func (d *Details) UnmarshalJSON(data []byte) error {
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

	d.InvalidMembers = invalidMembers

	for _, field := range memberFields {
		delete(tmp, field)
	}

	d.ExtensionMembers = tmp

	return nil
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
