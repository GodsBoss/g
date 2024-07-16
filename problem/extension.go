package problem

import (
	"encoding/json"
	"fmt"
)

// GetExtension attempts to create an instance of Extension that is populated with
// the fields from the details extension.
func GetExtension[Extension any](details Details) (*Extension, error) {
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

// SetExtension sets the extension members of details to the equivalent of the JSON representation of extension.
//
// Returns an error if extension cannot be marshalled into a JSON object.
func SetExtension(details *Details, extension any) error {
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
