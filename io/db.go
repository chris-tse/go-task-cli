package io

import (
	"encoding/json"
	"os"
)

func ParseJSONFile[T any](data []byte) ([]T, error) {
	var items []T
	err := json.Unmarshal(data, &items)

	// some error with parsing JSON
	if err != nil {
		return nil, err
	}

	return items, nil
}

func WriteToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
