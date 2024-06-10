// internal/filter/filter.go
package filter

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

// FilterJSONFile filters the JSON data in the input file based on the provided criteria and writes the filtered data to the output file.
func FilterJSONFile(inputFile, outputFile, criteria string) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    filteredData := make([]map[string]interface{}, 0)
    for _, item := range jsonData {
        if matchCriteria(item, criteria) {
            filteredData = append(filteredData, item)
        }
    }

    filteredJSON, err := json.MarshalIndent(filteredData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal filtered data: %w", err)
    }

    if err := ioutil.WriteFile(outputFile, filteredJSON, 0644); err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }

    return nil
}

// matchCriteria checks if the given item matches the criteria.
func matchCriteria(item map[string]interface{}, criteria string) bool {
    for _, value := range item {
        if fmt.Sprintf("%v", value) == criteria {
            return true
        }
    }
    return false
}
