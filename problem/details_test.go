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
		"default_type": func(t *testing.T, details problem.Details) {
			if typ := details.Type; typ != "about:blank" {
				t.Errorf("expected type '%s', got '%s'", "about:blank", typ)
			}
		},
		"detail": func(t *testing.T, details problem.Details) {
			if detail := details.Detail; detail != "This API does not understand SOAP." {
				t.Errorf("expected detail '%s', got '%s'", "This API does not understand SOAP.", detail)
			}
		},
		"extension": func(t *testing.T, details problem.Details) {
			extension := details.ExtensionMembers
			if extension == nil {
				t.Fatalf("expected non-nil extension")
			}
			if extensionType, ok := extension["type"]; ok {
				t.Errorf("expected no 'type' in extension, got %+v", extensionType)
			}
			errorList, ok := extension["errors"].([]interface{})
			if !ok {
				t.Error("expected 'errors' to contain array.")
			}
			if len(errorList) != 3 {
				t.Errorf("expected %d error list entries, got %+v", 3, errorList)
			}
		},
		"instance": func(t *testing.T, details problem.Details) {
			instance := details.Instance
			if instance != "https://api.example.org/problems/instances/666" {
				t.Errorf("expected instance '%s', got '%s'", "https://api.example.org/problems/instances/666", instance)
			}
		},
		"members_with_different_types": func(t *testing.T, details problem.Details) {
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
			invalidMembers := details.InvalidMembers
			if l := len(invalidMembers); l != 5 {
				t.Errorf("expected 5 invalid members, got %+v", invalidMembers)
			}
		},
		"status_400": func(t *testing.T, details problem.Details) {
			if status := details.Status; status != 400 {
				t.Errorf("expected status code %d, got %d", 400, status)
			}
			if expectedStatusText, actualStatusText := http.StatusText(400), details.StatusText(); expectedStatusText != actualStatusText {
				t.Errorf("expected status text '%s', got '%s'", expectedStatusText, actualStatusText)
			}
		},
		"status_too_big": func(t *testing.T, details problem.Details) {
			if status := details.Status; status != 0 {
				t.Errorf("expected status code %d, got %d", 0, status)
			}
		},
		"status_too_small": func(t *testing.T, details problem.Details) {
			if status := details.Status; status != 0 {
				t.Errorf("expected status code %d, got %d", 0, status)
			}
		},
		"title": func(t *testing.T, details problem.Details) {
			if title := details.Title; title != "JSON broken" {
				t.Errorf("expected title '%s', got '%s'", "JSON broken", title)
			}
		},
		"type": func(t *testing.T, details problem.Details) {
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

				var details problem.Details

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

	var details problem.Details

	if err := json.Unmarshal(data, &details); err == nil {
		t.Error("expected an error")
	}
}

func TestMarshalJSON(t *testing.T) {
	testcases := map[string]struct {
		details problem.Details
		check   func(t *testing.T, value map[string]interface{})
	}{
		"empty": {
			check: func(t *testing.T, values map[string]interface{}) {
				if len(values) != 1 {
					t.Errorf("expected only %d value, got %+v", 1, values)
				}
				if status, ok := values["status"]; ok {
					t.Errorf("expected no status, got %+v", status)
				}
				if title, ok := values["title"]; ok {
					t.Errorf("expected no title, got %+v", title)
				}
				if detail, ok := values["detail"]; ok {
					t.Errorf("expected no detail, got %+v", detail)
				}
				if instance, ok := values["instance"]; ok {
					t.Errorf("expected no instance, got %+v", instance)
				}
				typ, ok := values["type"]
				if !ok {
					t.Fatal("expected a type")
				}
				typAsString, ok := typ.(string)
				if !ok {
					t.Fatalf("expected type to be string, got %T", typ)
				}
				if typAsString != "about:blank" {
					t.Errorf("expected type to be '%s', got '%s'", "about:blank", typAsString)
				}
			},
		},
		"extension": {
			details: problem.Details{
				ExtensionMembers: map[string]any{
					"title": "ignored",
					"extra": "included",
				},
			},
			check: func(t *testing.T, values map[string]interface{}) {
				title, ok := values["title"]
				if ok {
					t.Errorf("expected no title, got %v", title)
				}
				extra, ok := values["extra"].(string)
				if !ok || extra != "included" {
					t.Errorf("expected 'extra' to be '%s' (string), got %v (%T)", "included", values["extra"], values["extra"])
				}
			},
		},
		"fields_populated": {
			details: problem.Details{
				Type:     "https://example.org/problems/not-enough-jquery",
				Status:   400,
				Title:    "Not enough jQuery",
				Detail:   "Your attempted solution does not contain enough jQuery to be accepted",
				Instance: "https://example.org/problems/not-enough-jquery#0001",
			},
			check: func(t *testing.T, values map[string]interface{}) {
				if len(values) != 5 {
					t.Errorf("expected %d values, got %+v", 5, values)
				}
				typ, ok := values["type"].(string)
				if !ok || typ != "https://example.org/problems/not-enough-jquery" {
					t.Errorf("expected 'type' to be '%s' (string), got %v (%T)", "https://example.org/problems/not-enough-jquery", values["type"], values["type"])
				}
				status, ok := values["status"].(float64)
				if !ok || int(status) != 400 {
					t.Errorf("expected 'status' to be %d (integer), got %v (%T)", 400, values["status"], values["status"])
				}
				title, ok := values["title"].(string)
				if !ok || title != "Not enough jQuery" {
					t.Errorf("expected 'title' to be '%s' (string), got %v (%T)", "Not enough jQuery", values["title"], values["title"])
				}
				detail, ok := values["detail"].(string)
				if !ok || detail != "Your attempted solution does not contain enough jQuery to be accepted" {
					t.Errorf("expected  'detail' to be '%s' (string), got %v (%T)", "Your attempted solution does not contain enough jQuery to be accepted", values["detail"], values["detail"])
				}
				instance, ok := values["instance"].(string)
				if !ok || instance != "https://example.org/problems/not-enough-jquery#0001" {
					t.Errorf("expected 'instance' to be '%s' (string), got %v (%T)", "https://example.org/problems/not-enough-jquery#0001", values["instance"], values["instance"])
				}
			},
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				data, err := json.Marshal(testcase.details)
				if err != nil {
					t.Fatalf("could not marshal problem details: %v", err)
				}

				var dest map[string]interface{}
				if err := json.Unmarshal(data, &dest); err != nil {
					t.Fatalf("could not unmarshal into map: %v", err)
				}

				testcase.check(t, dest)
			},
		)
	}
}

func TestGetTypeEmpty(t *testing.T) {
	var details problem.Details

	if typ := details.GetType(); typ != "about:blank" {
		t.Errorf("expected type to be '%s', got '%s'", "about:blank", typ)
	}
}

func TestGetTypeWithValue(t *testing.T) {
	expectedType := "http://example.com/problems/5423423424"
	details := problem.Details{
		Type: expectedType,
	}

	if typ := details.GetType(); typ != expectedType {
		t.Errorf("expected type to be '%s', got '%s'", typ, expectedType)
	}
}
