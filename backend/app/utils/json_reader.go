package utils

import (
	"encoding/json"
	"os"
)

func ReadJsonFile[T any](filename string) ([]T, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []T
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func WriteJsonFile[T any](filename string, data []T) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(data); err != nil {
		return err
	}
	return nil
}
