package optionaljson_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/GodsBoss/g/optional/encoding/optionaljson"
)

func TestErrorUnUnmarshalJSON(t *testing.T) {
	obj := struct {
		F optionaljson.Value[brokenJSON]
	}{
		F: optionaljson.NewValue(brokenJSON{}),
	}

	err := json.Unmarshal([]byte(`{ "f": "xyz" }`), &obj)

	if err == nil {
		t.Errorf("expected an error")
	}
}

type brokenJSON struct{}

func (obj *brokenJSON) Unmarshal(_ []byte) error {
	return errors.New("broken")
}
