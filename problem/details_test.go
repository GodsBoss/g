package problem_test

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/GodsBoss/g/problem"
)

func TestUnmarshalJSON(t *testing.T) {
	testcases := map[string]func(t *testing.T, details problem.Details){
		"status_400": func(t *testing.T, details problem.Details) {
			if status := details.Status; status != 400 {
				t.Errorf("expected status code %d, got %d", 400, status)
			}
			if expectedStatusText, actualStatusText := http.StatusText(400), details.StatusText(); expectedStatusText != actualStatusText {
				t.Errorf("expected status text '%s', got '%s'", expectedStatusText, actualStatusText)
			}
		},
		"title": func(t *testing.T, details problem.Details) {
			if title := details.Title; title != "JSON broken" {
				t.Errorf("expected title '%s', got '%s'", "JSON broken", title)
			}
		},
		"detail": func(t *testing.T, details problem.Details) {
			if detail := details.Detail; detail != "This API does not understand SOAP." {
				t.Errorf("expected detail '%s', got '%s'", "This API does not understand SOAP.", detail)
			}
		},
		"type": func(t *testing.T, details problem.Details) {
			typ := details.Type
			if typ != "https://api.example.org/foo/bar/123" {
				t.Errorf("expected type '%s', got '%s'", "https://api.example.org/foo/bar/123", typ)
			}
		},
		"instance": func(t *testing.T, details problem.Details) {
			instance := details.Instance
			if instance != "https://api.example.org/problems/instances/666" {
				t.Errorf("expected instance '%s', got '%s'", "https://api.example.org/problems/instances/666", instance)
			}
		},
	}

	for name := range testcases {
		check := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				filename := filepath.Join("testdata", name+".json")
				data, err := os.ReadFile(filename)
				if err != nil {
					t.Fatalf("could not read file %s: %v", filename, err)
				}

				var details problem.Details

				if err := json.Unmarshal(data, &details); err != nil {
					t.Fatalf("could not unmarshal payload: %v", err)
				}

				check(t, details)
			},
		)
	}
}
