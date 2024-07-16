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
