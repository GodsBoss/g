package problem

import "net/http"

type Details struct {
	// Status is the HTTP status code. May be 0.
	Status int
}

func (d Details) StatusText() string {
	return http.StatusText(d.Status)
}
