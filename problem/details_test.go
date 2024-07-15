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
	testcases := map[string]func(t *testing.T, details problem.Details[map[string]interface{}]){
		"default_type": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			if typ := details.Type; typ != "about:blank" {
				t.Errorf("expected type '%s', got '%s'", "about:blank", typ)
			}
		},
		"detail": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			if detail := details.Detail; detail != "This API does not understand SOAP." {
				t.Errorf("expected detail '%s', got '%s'", "This API does not understand SOAP.", detail)
			}
		},
		"extension": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			extension, err := details.Extension()
			if err != nil {
				t.Errorf("expected no extension error, got %v", err)
			}
			if extension == nil {
				t.Fatalf("expected non-nil extension")
			}
			if extensionType, ok := (*extension)["type"]; ok {
				t.Errorf("expected no 'type' in extension, got %+v", extensionType)
			}
			errorList, ok := (*extension)["errors"].([]interface{})
			if !ok {
				t.Error("expected 'errors' to contain array.")
			}
			if len(errorList) != 3 {
				t.Errorf("expected %d error list entries, got %+v", 3, errorList)
			}
		},
		"instance": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			instance := details.Instance
			if instance != "https://api.example.org/problems/instances/666" {
				t.Errorf("expected instance '%s', got '%s'", "https://api.example.org/problems/instances/666", instance)
			}
		},
		"members_with_different_types": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			if typ := details.Type; typ != "about:blank" {
				t.Errorf("expected type '%s', got '%s'", "about:blank", typ)
			}
			if status := details.Status; status != 0 {
				t.Errorf("expected status %d, got %d", 0, status)
			}
			if title := details.Title; title != "" {
				t.Errorf("expected no title, got '%s'", title)
			}
			if detail := details.Detail; detail != "" {
				t.Errorf("expected no detail, got '%s'", detail)
			}
			if instance := details.Instance; instance != "" {
				t.Errorf("expected no instance, got '%s'", instance)
			}
			invalidMembers := details.InvalidMembers()
			if l := len(invalidMembers); l != 5 {
				t.Errorf("expected 5 invalid members, got %+v", invalidMembers)
			}
		},
		"status_400": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			if status := details.Status; status != 400 {
				t.Errorf("expected status code %d, got %d", 400, status)
			}
			if expectedStatusText, actualStatusText := http.StatusText(400), details.StatusText(); expectedStatusText != actualStatusText {
				t.Errorf("expected status text '%s', got '%s'", expectedStatusText, actualStatusText)
			}
		},
		"title": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			if title := details.Title; title != "JSON broken" {
				t.Errorf("expected title '%s', got '%s'", "JSON broken", title)
			}
		},
		"type": func(t *testing.T, details problem.Details[map[string]interface{}]) {
			typ := details.Type
			if typ != "https://api.example.org/foo/bar/123" {
				t.Errorf("expected type '%s', got '%s'", "https://api.example.org/foo/bar/123", typ)
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

				var details problem.Details[map[string]interface{}]

				if err := json.Unmarshal(data, &details); err != nil {
					t.Fatalf("could not unmarshal payload: %v", err)
				}

				check(t, details)
			},
		)
	}
}

func TestUnmarshalMismatchingJSON(t *testing.T) {
	filename := "testdata/bool.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("could not read file '%s': %v", filename, err)
	}

	var details problem.Details[map[string]any]

	if err := json.Unmarshal(data, &details); err == nil {
		t.Error("expected an error")
	}
}
