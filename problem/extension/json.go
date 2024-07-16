// Package extension contains helpers for mapping extension fields as provided by problem details
// to and from concrete types.
package extension

import (
	"encoding/json"
	"fmt"

	"github.com/GodsBoss/g/problem"
)

// GetViaJSON attempts to create an instance of Extension that is populated with
// the fields from the details extension by converting the details extension fields
// into a JSON representatino which is then unmarshalled into an instance of Extension.
func GetViaJSON[Extension any](details problem.Details) (*Extension, error) {
	extensionMembers := details.ExtensionMembers
	if extensionMembers == nil {
		extensionMembers = make(map[string]any)
	}

	data, err := json.Marshal(extensionMembers)
	if err != nil {
		return nil, fmt.Errorf("could not marshal problem details extension members: %w", err)
	}

	var extension Extension
	if err := json.Unmarshal(data, &extension); err != nil {
		return nil, fmt.Errorf("could not unmarshal problem details extension members into *%T: %w", extension, err)
	}

	return &extension, nil
}

// SetViaJSON populates the details extension fields by creating an intermediate JSON object representation
// of the given extension.
func SetViaJSON(details *problem.Details, extension any) error {
	data, err := json.Marshal(extension)
	if err != nil {
		return fmt.Errorf("could not marshal extension: %w", err)
	}

	var extensionMembers map[string]any
	if err := json.Unmarshal(data, &extensionMembers); err != nil {
		return fmt.Errorf("could not unmarshal marshalled extension into map: %w", err)
	}

	details.ExtensionMembers = extensionMembers

	return nil
}
