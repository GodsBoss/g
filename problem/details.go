package problem

import (
	"encoding/json"
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
	type details Details

	var tmp details

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	d.Type = tmp.Type
	d.Status = tmp.Status
	d.Title = tmp.Title
	d.Detail = tmp.Detail
	d.Instance = tmp.Instance

	if d.Type == "" {
		d.Type = "about:blank"
	}

	return nil
}
