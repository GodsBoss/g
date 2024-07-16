package problem

import (
	"encoding/json"
	"fmt"
)

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
