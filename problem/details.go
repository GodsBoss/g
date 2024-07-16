package problem

import (
	"encoding/json"
	"math"
	"net/http"
)

// Details represents a problem details object that can be marshalled to and unmarshalled from JSON.
type Details struct {
	// Type is a URI reference that identifies the problem type. Defaults to "about:blank".
	Type string

	// Status is the HTTP status code.
	Status int

	// Title is a human-readable summary of the problem type.
	Title string

	// Detail is a human-readable explanation specific to this occurence of the problem.
	Detail string

	// Instance is a URI reference that identifies the specific occurence of the problem.
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

// GetType returns the details problem's type or "about:blank" if the type is empty.
func (d Details) GetType() string {
	if d.Type == "" {
		return "about:blank"
	}

	return d.Type
}

// StatusText returns the status text corresponding to the HTTP status code of this problem details.
// Returns an empty string for unknown status codes.
func (d Details) StatusText() string {
	return http.StatusText(d.Status)
}

// MarshalJSON marshals this Details object into a JSON representation. See [Details.ExtensionMembers] for
// handling of extension members.
func (d Details) MarshalJSON() ([]byte, error) {
	m := map[string]any{}

	for k, v := range d.ExtensionMembers {
		m[k] = v
	}

	for _, k := range memberFields {
		delete(m, k)
	}

	addNonZero(m, "type", d.Type)
	addNonZero(m, "status", d.Status)
	addNonZero(m, "title", d.Title)
	addNonZero(m, "detail", d.Detail)
	addNonZero(m, "instance", d.Instance)

	if d.Type == "" {
		m["type"] = "about:blank"
	}

	return json.Marshal(m)
}

// addNonZero adds an entry to a map if the value is not the zero value.
func addNonZero[T comparable](m map[string]any, key string, value T) {
	var zero T
	if value != zero {
		m[key] = value
	}
}

// UnmarshalJSON is very lenient. As long as the JSON data is an object, it can be unmarshalled into
// a problem details object. All members defined by RFC 9457 are optional.
//
// See [Details.InvalidMembers] for handling of members that share a name with the pre-defined members,
// but have a value with an incompatible type.
//
// See [Details.ExtensionMembers] for extension members.
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
