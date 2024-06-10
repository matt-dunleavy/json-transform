// internal/search/searcher.go
package search

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
)

// SearchJSONFile searches for specific values or patterns within a JSON file and writes the matching data to the output file.
func SearchJSONFile(inputFile, outputFile, pattern string) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    var results []map[string]interface{}
    for _, item := range jsonData {
        if matchesPattern(item, pattern) {
            results = append(results, item)
        }
    }

    resultsJSON, err := json.MarshalIndent(results, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal search results: %w", err)
    }

    if err := ioutil.WriteFile(outputFile, resultsJSON, 0644); err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }

    return nil
}

// matchesPattern checks if the given item matches the pattern.
func matchesPattern(item map[string]interface{}, pattern string) bool {
    // Convert pattern to lowercase for case-insensitive search
    pattern = strings.ToLower(pattern)
    for _, value := range item {
        if strings.Contains(strings.ToLower(fmt.Sprintf("%v", value)), pattern) {
            return true
        }
    }
    return false
}
