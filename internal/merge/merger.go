// internal/merge/merger.go
package merge

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "path/filepath"
)

// MergeJSONFiles merges multiple JSON files into a single JSON file.
func MergeJSONFiles(inputDir, outputFile string) error {
    var mergedData []map[string]interface{}

    files, err := ioutil.ReadDir(inputDir)
    if err != nil {
        return fmt.Errorf("failed to read input directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        filePath := filepath.Join(inputDir, file.Name())
        data, err := ioutil.ReadFile(filePath)
        if err != nil {
            return fmt.Errorf("failed to read file %s: %w", filePath, err)
        }

        var jsonData []map[string]interface{}
        err = json.Unmarshal(data, &jsonData)
        if err != nil {
            return fmt.Errorf("failed to unmarshal JSON data from file %s: %w", filePath, err)
        }

        mergedData = append(mergedData, jsonData...)
    }

    mergedJSON, err := json.MarshalIndent(mergedData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal merged data: %w", err)
    }

    if err := ioutil.WriteFile(outputFile, mergedJSON, 0644); err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }

    return nil
}
