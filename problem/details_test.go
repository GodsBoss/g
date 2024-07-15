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
