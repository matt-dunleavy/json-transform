// internal/split/splitter.go
package split

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "path/filepath"
)

// SplitJSONFile splits a large JSON file into smaller files with a specified number of objects per file.
func SplitJSONFile(inputFile, outputDir string, chunkSize int) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    totalObjects := len(jsonData)
    if chunkSize <= 0 {
        return fmt.Errorf("chunk size must be greater than zero")
    }

    for i := 0; i < totalObjects; i += chunkSize {
        end := i + chunkSize
        if end > totalObjects {
            end = totalObjects
        }

        chunk := jsonData[i:end]
        chunkData, err := json.MarshalIndent(chunk, "", "  ")
        if err != nil {
            return fmt.Errorf("failed to marshal chunk data: %w", err)
        }

        chunkFileName := filepath.Join(outputDir, fmt.Sprintf("chunk_%d.json", i/chunkSize))
        if err := ioutil.WriteFile(chunkFileName, chunkData, 0644); err != nil {
            return fmt.Errorf("failed to write chunk file: %w", err)
        }
    }

    return nil
}
