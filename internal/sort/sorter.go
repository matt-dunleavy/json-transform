// internal/sort/sorter.go
package sort

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "sort"
)

// SortJSONFile sorts JSON objects in the input file based on the specified field and writes the sorted data to the output file.
func SortJSONFile(inputFile, outputFile, field string) error {
    data, err := ioutil.ReadFile(inputFile)
    if err != nil {
        return fmt.Errorf("failed to read input file: %w", err)
    }

    var jsonData []map[string]interface{}
    if err := json.Unmarshal(data, &jsonData); err != nil {
        return fmt.Errorf("failed to unmarshal JSON data: %w", err)
    }

    sort.Slice(jsonData, func(i, j int) bool {
        return fmt.Sprintf("%v", jsonData[i][field]) < fmt.Sprintf("%v", jsonData[j][field])
    })

    sortedJSON, err := json.MarshalIndent(jsonData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal sorted data: %w", err)
    }

    if err := ioutil.WriteFile(outputFile, sortedJSON, 0644); err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }

    return nil
}
