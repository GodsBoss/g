package extension_test

import (
	"errors"
	"testing"

	"github.com/GodsBoss/g/problem"
	"github.com/GodsBoss/g/problem/extension"
)

func TestGetExtensionFailsDueToBrokenExtensionMembers(t *testing.T) {
	details := problem.Details{
		ExtensionMembers: map[string]any{
			"extra": jsonKiller{},
		},
	}

	extension, err := extension.GetViaJSON[map[string]any](details)
	if extension != nil {
		t.Errorf("expected no extension, got %+v", extension)
	}
	if err == nil {
		t.Error("expected an error")
	}
}

func TestGetExtensionFailsDueToBrokenTarget(t *testing.T) {
	details := problem.Details{}

	extension, err := extension.GetViaJSON[jsonKiller](details)
	if extension != nil {
		t.Errorf("expected no extension, got %+v", extension)
	}
	if err == nil {
		t.Error("expected an error")
	}
}

func TestSetExtensionFailsDueToUnmarshallableExtension(t *testing.T) {
	details := problem.Details{}
	err := extension.SetViaJSON(&details, jsonKiller{})
	if err == nil {
		t.Errorf("expected an error")
	}
}

func TestSetExtensionFailsDueToExtensionThatDoesNotMarshalIntoMap(t *testing.T) {
	details := problem.Details{}
	err := extension.SetViaJSON(&details, true)
	if err == nil {
		t.Errorf("expected an error")
	}
}

type jsonKiller struct{}

func (jsonKiller) MarshalJSON() ([]byte, error) {
	return nil, errors.New("can't marshal this")
}

func (*jsonKiller) UnmarshalJSON(_ []byte) error {
	return errors.New("can't unmarshal this")
}
