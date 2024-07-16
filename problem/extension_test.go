package problem_test

import (
	"errors"
	"testing"

	"github.com/GodsBoss/g/problem"
)

func TestGetExtensionFailsDueToBrokenExtensionMembers(t *testing.T) {
	details := problem.Details{
		ExtensionMembers: map[string]any{
			"extra": jsonKiller{},
		},
	}

	extension, err := problem.GetExtension[map[string]any](details)
	if extension != nil {
		t.Errorf("expected no extension, got %+v", extension)
	}
	if err == nil {
		t.Error("expected an error")
	}
}

func TestGetExtensionFailsDueToBrokenTarget(t *testing.T) {
	details := problem.Details{}

	extension, err := problem.GetExtension[jsonKiller](details)
	if extension != nil {
		t.Errorf("expected no extension, got %+v", extension)
	}
	if err == nil {
		t.Error("expected an error")
	}
}

type jsonKiller struct{}

func (jsonKiller) MarshalJSON() ([]byte, error) {
	return nil, errors.New("can't marshal this")
}

func (*jsonKiller) UnmarshalJSON(_ []byte) error {
	return errors.New("can't unmarshal this")
}
