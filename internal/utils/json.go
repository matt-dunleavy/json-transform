// internal/utils/json.go
package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

// ReadJSONFile reads and unmarshals a JSON file into a provided interface.
func ReadJSONFile(filePath string, v interface{}) error {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to read file %s: %w", filePath, err)
    }

    if err := json.Unmarshal(data, v); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data from file %s: %w", filePath, err)
    }

    return nil
}

// WriteJSONFile marshals a provided interface and writes it to a JSON file.
func WriteJSONFile(filePath string, v interface{}) error {
    data, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal data to JSON for file %s: %w", filePath, err)
    }

    if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
        return fmt.Errorf("failed to write JSON data to file %s: %w", filePath, err)
    }

    return nil
}

// PrettyPrintJSON returns a pretty-printed JSON string from a provided interface.
func PrettyPrintJSON(v interface{}) (string, error) {
    data, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return "", fmt.Errorf("failed to marshal data to JSON: %w", err)
    }

    return string(data), nil
}
