package problem

import "net/http"

type Details struct {
	// Status is the HTTP status code. May be 0.
	Status int

	// Title is a human-readable summary of the problem type.
	Title string
}

func (d Details) StatusText() string {
	return http.StatusText(d.Status)
}
